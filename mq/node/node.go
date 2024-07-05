package node

import (
	"dominant/message"
	"dominant/mq"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 16
//

type Node struct {
	Addr      string           //节点地址
	MQ        *mq.MessageQueue //专有消息队列
	NodeType  string           //节点类型
	TopicList []string         //节点订阅主题列表
}

var defaultChanSize = 100

func NewNode(ip, nType string) *Node {
	//id := uuid.NewV4().String()
	return &Node{
		Addr: ip,
		MQ: &mq.MessageQueue{
			MessageHistory: make(map[string]*message.Message),
			MessageChan:    make(chan *message.Message, defaultChanSize),
		},
		NodeType:  nType,
		TopicList: []string{},
	}
}
