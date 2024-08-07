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
	Content     any      `json:"content"`
}

func NewMessage(src string, dstList []string, content any) *Message {
	id := ""
	return &Message{
		ID:          id,
		Src:         src,
		DstList:     dstList,
		CreateTime:  time.Now().Format(dateTimeFormat),
		ConsumeTime: "",
		Content:     content,
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

//func (m *Distribute) GetContentFromJson(bytesMessage []byte) string {
//	err := m.MessageJsonUnMarshal(bytesMessage)
//	if err != nil {
//		return ""
//	}
//	return m.Content.(string)
//}
