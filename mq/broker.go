package mq

import (
	"dominant/message"
	"dominant/mq/node"
	"sync"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

type nodeMap map[string]*node.Node
type messageChanSlice []chan *message.Message
type MessageMap map[string]*message.Message //全局消息map

type broker struct {
	OnlineNodes nodeMap               //所有在线节点
	Subscribers map[string]nodeMap    //节点主题订阅关系组
	Messages    MessageMap            //所有消息
	MessagesMap map[string]MessageMap //消息主题关系组
	MessageChan chan *message.Message
	rwm         sync.RWMutex
}

// Distribute 消息分发
func (b *broker) Distribute(msg *message.Message) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	if msg.Dst != "" {

	}
	if nm, ok := b.Subscribers[msg.Topic]; ok {
		var mChans messageChanSlice
		for _, v := range nm {
			mChans = append(mChans, v.SendChan)
		}
		//启动一个协程将消息发送到各个通道
		go func(msg *message.Message, mcs messageChanSlice) {
			for _, c := range mcs {
				c <- msg
			}
		}(msg, mChans)
	}
}

// Subscribe 为某个ip订阅主题
func (b *broker) Subscribe(ip string, topics []string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	n := b.OnlineNodes[ip]
	for _, topic := range topics {
		if current, ok := b.Subscribers[topic]; ok {
			current[ip] = n
		} else {
			//当前主题不存在节点map创建新的map
			nm := make(nodeMap)
			nm[ip] = n
			b.Subscribers[topic] = nm
		}
	}
}

// Unsubscribe 为某个ip取消订阅主题
func (b *broker) Unsubscribe(ip string, topics []string) {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	for _, topic := range topics {
		if current, ok := b.Subscribers[topic]; ok {
			delete(current, ip)
		}
	}
}

// ListSubscribers 列出所有订阅者
func (b *broker) ListSubscribers() []string {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var res []string
	for _, sb := range b.Subscribers {
		for _, n := range sb {
			res = append(res, n.Addr)
		}
	}
	return res
}
