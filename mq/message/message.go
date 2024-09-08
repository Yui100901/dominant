package message

import (
	"dominant/mq"
	"math/rand"
	"strconv"
	"time"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 17
//

type Message struct {
	ID            string   `json:"id"`
	Topic         string   `json:"topic"`
	Type          string   `json:"type"`
	Src           string   `json:"src"`           //消息来源,某个节点的id
	PresetDstList []string `json:"presetDstList"` //预设消息目的地列表
	ActualDstList []string `json:"actualDstList"` //实际消息目的地列表
	CreateTime    string   `json:"createTime"`
	ConsumeTime   string   `json:"consumeTime"`
	Content       any      `json:"content"`
}

func NewMessage(topic, messageType, src string, dstList []string, content any) *Message {
	id := strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	return &Message{
		ID:            id,
		Topic:         topic,
		Type:          messageType,
		Src:           src,
		PresetDstList: dstList,
		ActualDstList: []string{},
		CreateTime:    time.Now().Format(mq.DateTimeFormat),
		ConsumeTime:   "",
		Content:       content,
	}
}

// AddPresetDestination 增加预设目的地
func (m *Message) AddPresetDestination(dst string) {
	m.PresetDstList = append(m.PresetDstList, dst)
}

//func (m *Distribute) GetContentFromJson(bytesMessage []byte) string {
//	err := m.MessageJsonUnMarshal(bytesMessage)
//	if err != nil {
//		return ""
//	}
//	return m.Content.(string)
//}
