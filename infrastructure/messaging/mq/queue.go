package mq

import (
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/5 09 25
//

type Queue struct {
	MessageHistory map[string]*Message //消息历史
	MessageChan    chan *Message       //消息通道，用于发送消息
	SaveHistory    bool
	rwm            sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		MessageHistory: make(map[string]*Message),
		SaveHistory:    false,
		MessageChan:    make(chan *Message, DefaultChanSize),
	}
}

func (mq *Queue) Enqueue(msg *Message) {
	mq.rwm.Lock()
	defer mq.rwm.Unlock()
	if mq.SaveHistory {
		mq.MessageHistory[msg.ID] = msg
	}
	mq.MessageChan <- msg
}

func (mq *Queue) Dequeue() *Message {
	select {
	case msg := <-mq.MessageChan:
		msg.ConsumeTime = time.Now().Format(DateTimeFormat)
		return msg
	default:
		return &Message{}
	}
}
