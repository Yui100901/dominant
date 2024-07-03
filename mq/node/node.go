package node

import "dominant/message"

//
// @Author yfy2001
// @Date 2024/6/19 21 16
//

type Node struct {
	Addr        string                //节点地址
	SendChan    chan *message.Message //节点发送通道
	ReceiveChan chan *message.Message //节点接收通道
	NodeType    string                //节点类型
	TopicList   []string              //节点订阅主题列表
}

var defaultChanSize = 100

func NewNode(ip, nType string) *Node {
	//id := uuid.NewV4().String()
	return &Node{
		Addr:        ip,
		SendChan:    make(chan *message.Message, defaultChanSize),
		ReceiveChan: make(chan *message.Message, defaultChanSize),
		NodeType:    nType,
		TopicList:   []string{},
	}
}
