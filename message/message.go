package message

import (
	"encoding/json"
	"time"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 17
//

const dateTimeFormat = "2006-01-02 15:04:05"

type Message struct {
	ID          string    `json:"id"`
	Dst         string    `json:"dst"`       //目的地
	Topic       string    `json:"topic"`     //消息主题
	Broadcast   bool      `json:"broadcast"` //是否广播
	CreateTime  time.Time `json:"createTime"`
	ConsumeTime time.Time `json:"consumeTime"`
	Body        any       `json:"body"`
}

func (m *Message) MessageJsonMarshal() ([]byte, error) {
	bytesMessage, err := json.Marshal(*m)
	if err != nil {
		return nil, err
	}
	return bytesMessage, nil
}

func (m *Message) MessageJsonUnMarshal(bytesMessage []byte) error {
	err := json.Unmarshal(bytesMessage, m)
	if err != nil {
		return err
	}
	return nil
}

func NewMessage(dst, topic, content string) *Message {
	id := ""
	return &Message{
		ID:          id,
		Dst:         dst,
		Topic:       topic,
		Broadcast:   false,
		CreateTime:  time.Now(),
		ConsumeTime: time.Time{},
		Body:        content,
	}
}
