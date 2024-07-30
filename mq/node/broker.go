package node

import (
	"dominant/mq/message"
	"log"
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
	OnlineNodeMap map[string]*Node          //所有在线节点
	MQMap         map[string]*message.Queue //主题消息队列
	MessageChan   chan *message.Message
	rwm           sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		//OnlineNodes: make(nodeMap),
		OnlineNodeMap: make(map[string]*Node),
		MQMap:         make(map[string]*message.Queue),
		rwm:           sync.RWMutex{},
	}
}

// Distribute 消息分发
func (b *Broker) Distribute(msg *message.Message) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	//发送到每一个目的地节点的通道
	for _, dst := range msg.DstList {
		if n, ok := b.OnlineNodeMap[dst]; ok {
			n.MQ.Enqueue(msg)
		}
	}
}

// Register 将某个id注册为在线节点
func (b *Broker) Register(id string, ip string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	var n *Node
	n = b.OnlineNodeMap[id]
	if n == nil {
		//如果id为空则向全局队列中注册
		n = NewNode(id, ip)
		b.OnlineNodeMap[id] = n
		//启动保活协程
		go b.keepAlive(id)
	} else {
		//如果id已经存在则向目标节点发送保活消息
		n.AliveChan <- true
	}

}

// Unregister 取消某个id的节点在线状态
func (b *Broker) Unregister(id string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	delete(b.OnlineNodeMap, id)
}

// GetNodeById 根据id获取一个节点
func (b *Broker) GetNodeById(id string) *Node {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	n := b.OnlineNodeMap[id]
	return n
}

// ListNodes 列出所有在线节点
func (b *Broker) ListNodes() []string {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var list []string
	for _, n := range b.OnlineNodeMap {
		list = append(list, n.ID)
	}
	return list
}

// GetMessage 根据id定位一则消息
func (b *Broker) GetMessage(id string) *message.Message {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	msg := &message.Message{}
	if n, ok := b.OnlineNodeMap[id]; ok {
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
