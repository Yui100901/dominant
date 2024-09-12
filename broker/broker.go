package broker

import (
	"context"
	"dominant/mq"
	"dominant/redis_utils"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

type nodeMap map[string]*Node
type messageChanSlice []chan *mq.Message

type Broker struct {
	//OnlineNodes nodeMap
	NodeMap map[string]*Node //所有节点
	MainMQ  *mq.Queue        //全局主队列
	rwm     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		//OnlineNodes: make(nodeMap),
		NodeMap: make(map[string]*Node),
		MainMQ:  mq.NewQueue(),
		rwm:     sync.RWMutex{},
	}
}

// Distribute 消息分发实现，根据预设目的地设置实际目的地并进行分发
func (b *Broker) Distribute() <-chan *mq.Message {
	for {
		select {
		case msg := <-b.MainMQ.MessageChan:
			//获取当前在线节点列表
			nodeIDs, nodeList := b.GetAliveNodeIDList()
			//消息有主题则根据主题增加预设目的地
			if msg.Topic != "" {
				for _, n := range nodeList {
					if _, ok := n.TopicMap[""]; ok {
						msg.AddPresetDestination(n.ID)
					}
				}
			}
			if len(msg.PresetDstList) != 0 {
				//将预设目的地设置为实际目的地
				msg.ActualDstList = msg.PresetDstList
			} else {
				//当消息预设目的地为空时将随机分配消息目的地
				dst := randomStringFromSlice(nodeIDs)
				msg.ActualDstList = append(msg.ActualDstList, dst)
			}
			log.Println("Distribute:", msg.ActualDstList)
			go b.Send(msg)
		}
	}
}

func randomStringFromSlice(slice []string) string {
	rand.NewSource(time.Now().UnixNano()) // 设置随机数种子
	return slice[rand.Intn(len(slice))]
}

// Send 消息发送
func (b *Broker) Send(msg *mq.Message) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	if msg.Type == "command" {
		//发送到每一个目的地节点的通道
		for _, dst := range msg.ActualDstList {
			if n, ok := b.NodeMap[dst]; ok {
				n.MQ.Enqueue(msg)
			}
		}
	}
}

// Login 节点登录
func (b *Broker) Login(id, addr string, state []byte) string {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	n := b.NodeMap[id]
	if n != nil {
		//id已经存在
		if n.IsAlive {
			//该节点存活，使存活节点下线
			b.Unregister(id)
		}
	}
	token := uuid.NewString()
	//重新给节点分配token
	n = NewNode(id, addr, token, state)
	n.RealtimeInfo = state
	//存入全局节点表
	b.NodeMap[id] = n
	log.Println("登录id", id)
	log.Printf("%v", b.NodeMap[id].Token)
	//启动保活协程
	go b.keepAlive(id)
	ctx := context.Background()
	//刷新redis存储的最新状态
	redis_utils.GlobalRedisClient.Set(ctx, n.ID, state, 60*time.Second)
	return token
}

// Verify 节点在线验证
func (b *Broker) Verify(id, token string, state []byte) bool {
	b.rwm.Lock()
	defer b.rwm.Unlock()
	n := b.NodeMap[id]
	if n == nil {
		//节点未登录
		log.Println("非法的节点id", id)
		return false
	} else {
		//id已经存在
		if n.IsAlive {
			//该节点存活，向目标节点发送保活消息
			if n.Token == token {
				n.AliveChan <- true
			} else {
				log.Println("token过期")
				return false
			}
		} else {
			log.Println("请重新登录")
			return false
		}
		n.RealtimeInfo = state
	}
	ctx := context.Background()
	//刷新redis存储的最新状态
	redis_utils.GlobalRedisClient.Set(ctx, n.ID, state, 60*time.Second)
	return true
}

// Unregister 取消某个id的节点在线状态
func (b *Broker) Unregister(id string) {
	b.NodeMap[id].AliveChan <- false
}

// GetNodeById 根据id获取一个节点
func (b *Broker) GetNodeById(id string) *Node {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	n := b.NodeMap[id]
	return n
}

// GetAliveNodeIDList 获取所有在线节点ID和节点详细信息
func (b *Broker) GetAliveNodeIDList() ([]string, []Node) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var onlineNodeIdList []string
	var onlineNodeList []Node
	for id, n := range b.NodeMap {
		if n.IsAlive {
			onlineNodeIdList = append(onlineNodeIdList, id)
			onlineNodeList = append(onlineNodeList, *n)
		}
	}
	return onlineNodeIdList, onlineNodeList
}

// GetAliveNodeMessage 获取所有节点最新状态消息
func (b *Broker) GetAliveNodeMessage() []any {
	idList, _ := b.GetAliveNodeIDList()
	fmt.Println("Online Node List:", idList)
	ctx := context.Background()
	messageList, err := redis_utils.GlobalRedisClient.MGet(ctx, idList...).Result()
	if err != nil {
		fmt.Println("Get Alive Node Message Error:", err)
	}
	return messageList
}

// GetMessage 根据节点id定位一则消息
func (b *Broker) GetMessage(nodeId string) *mq.Message {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	msg := &mq.Message{}
	if n, ok := b.NodeMap[nodeId]; ok {
		msg = n.MQ.Dequeue()
	}
	return msg
}

// KeepAlive 保持在线
func (b *Broker) keepAlive(id string) {
	n := b.GetNodeById(id)
	for {
		select {
		case alive := <-n.AliveChan:
			if alive {
				continue
			} else {
				//收到下线指令，退出保活协程
				n.IsAlive = false
				return
			}
		case <-time.After(time.Second * 60):
			log.Println(id, "超时退出！")
			n.IsAlive = false
			b.Unregister(id)
		}
	}
}
