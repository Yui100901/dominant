package message

import (
	"encoding/json"
	"time"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 17
//

type Message struct {
	ID          string   `json:"id"`
	Src         string   `json:"src"` //消息来源
	DstList     []string `json:"dst"` //消息目的地列表
	CreateTime  string   `json:"createTime"`
	ConsumeTime string   `json:"consumeTime"`
	Body        any      `json:"body"`
}

func NewMessage(src string, dst []string, content any) *Message {
	id := ""
	return &Message{
		ID:          id,
		Src:         src,
		DstList:     dst,
		CreateTime:  time.Now().Format(dateTimeFormat),
		ConsumeTime: "",
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
