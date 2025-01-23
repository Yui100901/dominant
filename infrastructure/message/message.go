package message

import (
	"github.com/google/uuid"
	"time"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 17
//

type Message struct {
	ID         string    `json:"id"`
	NodeId     string    `json:"nodeId"`
	CreateTime time.Time `json:"createTime"`
	Content    any       `json:"content"`
}

func NewMessage(nodeId string, content any) *Message {
	id := uuid.NewString()
	return &Message{
		ID:         id,
		NodeId:     nodeId,
		CreateTime: time.Now(),
		Content:    content,
	}
}

//type MessageConverter interface {
//	ConvertToMessage() *Message[]
//}
