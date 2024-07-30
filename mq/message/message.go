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
	Src         string    `json:"src"` //消息来源
	DstList     []string  `json:"dst"` //消息目的地列表
	CreateTime  time.Time `json:"createTime"`
	ConsumeTime time.Time `json:"consumeTime"`
	Body        any       `json:"body"`
}

func NewMessage(src string, dst []string, content any) *Message {
	id := ""
	return &Message{
		ID:          id,
		Src:         src,
		DstList:     dst,
		CreateTime:  time.Now(),
		ConsumeTime: time.Time{},
		Body:        content,
	}
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
