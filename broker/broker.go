package broker

import (
	"context"
	"dominant/mq/message"
	"dominant/mq/node"
	"dominant/redis_utils"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

type nodeMap map[string]*node.Node
type messageChanSlice []chan *message.Message

type Broker struct {
	//OnlineNodes nodeMap
	NodeMap map[string]*node.Node //所有节点
	MainMQ  *message.Queue        //全局主队列
	rwm     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		//OnlineNodes: make(nodeMap),
		NodeMap: make(map[string]*node.Node),
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
			nodeIDs, nodeList := b.GetAliveNodeIDList()
			//消息有主题则根据主题增加预设目的地
			if msg.Topic != "" {
				for _, n := range nodeList {
					if _, ok := n.TopicMap[""]; ok {
						msg.AddPresetDestination(n.ID)
					}
				}
			}
			if len(msg.PresetDstList) != 0 {
				//将预设目的地设置为实际目的地
				msg.ActualDstList = msg.PresetDstList
			} else {
				//当消息预设目的地为空时将随机分配消息目的地
				dst := randomStringFromSlice(nodeIDs)
				msg.ActualDstList = append(msg.ActualDstList, dst)
			}
			log.Println("Distribute:", msg.ActualDstList)
			go b.Send(msg)
		}
	}
}

func randomStringFromSlice(slice []string) string {
	rand.NewSource(time.Now().UnixNano()) // 设置随机数种子
	return slice[rand.Intn(len(slice))]
}

// Send 消息发送
func (b *Broker) Send(msg *message.Message) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	if msg.Type == "command" {
		//发送到每一个目的地节点的通道
		for _, dst := range msg.ActualDstList {
			if n, ok := b.NodeMap[dst]; ok {
				n.MQ.Enqueue(msg)
			}
		}
	}
}

// Register 将某个id注册为在线节点
func (b *Broker) Register(id, addr string, state []byte) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	n := b.NodeMap[id]
	if n == nil {
		//id为空则向全局map中注册
		n = node.NewNode(id, addr, state)
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
		n.RealtimeInfo = state
	}
	ctx := context.Background()
	redis_utils.GlobalRedisClient.Set(ctx, n.ID, state, 60*time.Second)
}

// Unregister 取消某个id的节点在线状态
func (b *Broker) Unregister(id string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	b.NodeMap[id].IsAlive = false
}

// GetNodeById 根据id获取一个节点
func (b *Broker) GetNodeById(id string) *node.Node {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	n := b.NodeMap[id]
	return n
}

// GetAliveNodeIDList 获取所有在线节点ID和节点详细信息
func (b *Broker) GetAliveNodeIDList() ([]string, []node.Node) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var onlineNodeIdList []string
	var onlineNodeList []node.Node
	for id, n := range b.NodeMap {
		if n.IsAlive {
			onlineNodeIdList = append(onlineNodeIdList, id)
			onlineNodeList = append(onlineNodeList, *n)
		}
	}
	return onlineNodeIdList, onlineNodeList
}

// GetAliveNodeMessage 获取所有节点最新状态消息
func (b *Broker) GetAliveNodeMessage() []any {
	idList, _ := b.GetAliveNodeIDList()
	fmt.Println("Online Node List:", idList)
	ctx := context.Background()
	messageList, err := redis_utils.GlobalRedisClient.MGet(ctx, idList...).Result()
	if err != nil {
		fmt.Println("Get Alive Node Message Error:", err)
	}
	return messageList
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
