package node

import (
	"dominant/mq/message"
	"log"
	"math/rand"
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

type nodeMap map[string]*Node
type messageChanSlice []chan *message.Message

type Broker struct {
	//OnlineNodes nodeMap
	NodeMap map[string]*Node //所有节点
	MainMQ  *message.Queue   //全局队列
	rwm     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		//OnlineNodes: make(nodeMap),
		NodeMap: make(map[string]*Node),
		MainMQ:  message.NewMessageQueue(),
		rwm:     sync.RWMutex{},
	}
}

// Distribute 消息分发实现，根据预设目的地设置实际目的地并进行分发
func (b *Broker) Distribute() <-chan *message.Message {
	for {
		select {
		case msg := <-b.MainMQ.MessageChan:
			//获取当前在线节点列表
			nodes := b.GetAliveNodeIDList()
			if msg.PresetDstList == nil {
				//当消息预设目的地为空时将随机分配消息目的地
				dst := randomStringFromSlice(nodes)
				msg.ActualDstList = append(msg.ActualDstList, dst)
			} else {
				//将预设目的地设置为实际目的地
				msg.ActualDstList = msg.PresetDstList
			}
			go b.Send(msg)
		}
	}
}

func randomStringFromSlice(slice []string) string {
	rand.NewSource(time.Now().UnixNano()) // 设置随机数种子
	return slice[rand.Intn(len(slice))]
}

// Send 消息分发
func (b *Broker) Send(msg *message.Message) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	//发送到每一个目的地节点的通道
	for _, dst := range msg.ActualDstList {
		if n, ok := b.NodeMap[dst]; ok {
			n.MQ.Enqueue(msg)
		}
	}
}

// Register 将某个id注册为在线节点
func (b *Broker) Register(id string, ip string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	n := b.NodeMap[id]
	if n == nil {
		//id为空则向全局map中注册
		n = NewNode(id, ip)
		b.NodeMap[id] = n
		//启动保活协程
		go b.keepAlive(id)
	} else {
		//id已经存在
		if n.IsAlive {
			//该节点存活，向目标节点发送保活消息
			n.AliveChan <- true
		} else {
			//该节点未存活，则使该节点重新上线
			n.IsAlive = true
			go b.keepAlive(id)
		}
	}

}

// Unregister 取消某个id的节点在线状态
func (b *Broker) Unregister(id string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	b.NodeMap[id].IsAlive = false
}

// GetNodeById 根据id获取一个节点
func (b *Broker) GetNodeById(id string) *Node {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	n := b.NodeMap[id]
	return n
}

// GetAliveNodeIDList 获取所有在线节点ID
func (b *Broker) GetAliveNodeIDList() []string {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var list []string
	for _, n := range b.NodeMap {
		if n.IsAlive {
			list = append(list, n.ID)
		}
	}
	return list
}

// GetMessage 根据id定位一则消息
func (b *Broker) GetMessage(id string) *message.Message {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	msg := &message.Message{}
	if n, ok := b.NodeMap[id]; ok {
		msg = n.MQ.Dequeue()
	}
	return msg
}

// KeepAlive 保持在线
func (b *Broker) keepAlive(id string) {
	n := b.GetNodeById(id)
	for {
		select {
		case <-n.AliveChan:
			continue
		case <-time.After(time.Second * 60):
			b.Unregister(id)
			log.Println(id, "超时退出！")
			return // 超时，说明未在线
		}
	}
}
