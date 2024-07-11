package message

import (
	"sync"
)

//
// @Author yfy2001
// @Date 2024/7/5 09 25
//

type MessageQueue struct {
	MessageHistory map[string]*Message //消息历史
	MessageChan    chan *Message       //消息通道，用于发送消息
	rwm            sync.RWMutex
}

var defaultChanSize = 100

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		MessageHistory: make(map[string]*Message),
		MessageChan:    make(chan *Message, defaultChanSize),
	}
}

func (mq *MessageQueue) Enqueue(msg *Message) {
	mq.rwm.Lock()
	defer mq.rwm.Unlock()
	mq.MessageHistory[msg.ID] = msg
	mq.MessageChan <- msg
}

func (mq *MessageQueue) Dequeue() *Message {
	select {
	case msg := <-mq.MessageChan:
		return msg
	default:
		return &Message{}
	}
}
