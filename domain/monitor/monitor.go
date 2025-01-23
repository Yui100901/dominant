package monitor

import (
	"github.com/Yui100901/MyGo/concurrent"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

// Monitor 监视器，用于监测每个节点的状态
type Monitor struct {
	NodeMap    *concurrent.SafeMap[string, *Node] //节点map
	NodeIdChan chan string
	CleanTime  time.Duration
}

func NewMonitor(cleanTime time.Duration) *Monitor {
	return &Monitor{
		NodeMap:    concurrent.NewSafeMap[string, *Node](32),
		NodeIdChan: make(chan string),
		CleanTime:  cleanTime,
	}
}

func (m *Monitor) Start() {
	for {
		select {
		case nodeId := <-m.NodeIdChan:
			MainMonitor.Connect(nodeId)
		}
	}
}

func (m *Monitor) Clean() {
	for {
		select {
		case <-time.After(m.CleanTime):
			offlineNodeIds := m.GetOfflineNodes()
			for _, id := range offlineNodeIds {
				m.Disconnect(id)
			}
		}
	}
}

// Connect 连接函数，刷新某个节点的状态
func (m *Monitor) Connect(nodeId string) {
	var specificNode *Node
	if n, ok := m.NodeMap.Get(nodeId); ok {
		specificNode = n
	} else {
		specificNode = NewNode(nodeId)
		m.NodeMap.Set(nodeId, specificNode)
		go specificNode.AliveCheck()
	}
	specificNode.AliveChan <- struct{}{}
}

func (m *Monitor) Disconnect(nodeId string) {
	node, _ := m.NodeMap.Get(nodeId)
	node.Exit()
	m.NodeMap.Delete(nodeId)
}

// GetOnlineNodes 获取所有在线节点
func (m *Monitor) GetOnlineNodes() []string {
	var onlineNodes []string
	m.NodeMap.ForEach(func(k string, node *Node) bool {
		if node != nil && node.IsAlive {
			onlineNodes = append(onlineNodes, k)
		}
		return true
	})
	return onlineNodes
}

// GetOfflineNodes 获取所有未在线节点
func (m *Monitor) GetOfflineNodes() []string {
	var offlineNodes []string
	m.NodeMap.ForEach(func(k string, node *Node) bool {
		if node != nil && !node.IsAlive {
			offlineNodes = append(offlineNodes, k)
		}
		return true
	})
	return offlineNodes
}
