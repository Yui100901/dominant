package node

import (
	"dominant/domain/node/authentication"
	"dominant/infrastructure/messaging/mq"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 16
//

type Node struct {
	ID        string
	Addr      string    //节点地址
	MQ        *mq.Queue //节点队列
	IsAlive   bool      //存活标记
	AliveChan chan bool //心跳通道
	//NodeType     string            //节点类型
	Auth         *authentication.Authentication
	TopicMap     map[string]string //节点订阅主题表
	RealtimeInfo any               //实时数据
}

func NewNode(id, addr string, info any) *Node {
	//id := uuid.NewV4().String()
	return &Node{
		ID:           id,
		Addr:         addr,
		MQ:           mq.NewQueue(),
		IsAlive:      true, //节点创建时默认存活状态
		AliveChan:    make(chan bool),
		Auth:         authentication.NewAuthentication(id, 0),
		TopicMap:     make(map[string]string),
		RealtimeInfo: info,
	}
}
