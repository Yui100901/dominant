package node

import (
	"dominant/mq/message"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 16
//

type Node struct {
	ID        string         //节点唯一标识符
	Addr      string         //节点地址
	MQ        *message.Queue //专有消息队列
	AliveChan chan bool
	//NodeType string                //节点类型
	//TopicList []string              //节点订阅主题列表
}

func NewNode(id, ip string) *Node {
	//id := uuid.NewV4().String()

	return &Node{
		ID:        id,
		Addr:      ip,
		MQ:        message.NewMessageQueue(),
		AliveChan: make(chan bool),
	}
}
