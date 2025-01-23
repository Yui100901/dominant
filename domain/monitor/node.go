package monitor

import (
	"github.com/Yui100901/MyGo/log_utils"
	"github.com/google/uuid"
	"time"
)

//
// @Author yfy2001
// @Date 2024/6/19 21 16
//

type Node struct {
	ID            string
	AliveChan     chan struct{} //心跳通道
	Done          chan struct{} //退出信号通道
	IsAlive       bool          //存活标记
	LastAliveTime time.Time     //最后存活时间
}

func NewNode(id string) *Node {
	if id == "" {
		id = uuid.NewString()
	}
	return &Node{
		ID:        id,
		AliveChan: make(chan struct{}),
		IsAlive:   false,
	}
}

func (n *Node) AliveCheck() {
	for {
		select {
		case <-n.AliveChan:
			if !n.IsAlive {
				log_utils.Info.Println(n.ID, "上线！")
			}
			n.IsAlive = true
			n.LastAliveTime = time.Now()
		case <-n.Done:
			return
		case <-time.After(10 * time.Second):
			if n.IsAlive {
				log_utils.Warn.Println(n.ID, "超时下线！")
			}
			n.IsAlive = false
		}
	}
}

func (n *Node) Exit() {
	n.Done <- struct{}{}
}
