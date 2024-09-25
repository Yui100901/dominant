package broker

import (
	"context"
	"dominant/domain/node"
	"dominant/infrastructure/messaging/mq"
	"dominant/infrastructure/utils/log_utils"
	"dominant/infrastructure/utils/network/redis_utils"
	"dominant/infrastructure/utils/rand_utils"
	"github.com/google/uuid"
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/7/3 13 49
//

type nodeMap map[string]*node.Node
type messageChanSlice []chan *mq.Message

type Broker struct {
	//OnlineNodes nodeMap
	NodeMap map[string]*node.Node //所有节点
	MainMQ  *mq.Queue             //全局主队列
	rwm     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		//OnlineNodes: make(nodeMap),
		NodeMap: make(map[string]*node.Node),
		MainMQ:  mq.NewQueue(),
		rwm:     sync.RWMutex{},
	}
}

// GetNodeById 根据id获取一个节点
func (b *Broker) GetNodeById(id string) *node.Node {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	n := b.NodeMap[id]
	return n
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
				dst := rand_utils.RandomFromSlice[string](nodeIDs)
				msg.ActualDstList = append(msg.ActualDstList, dst)
			}
			log_utils.Info.Println("Distribute:", msg.ActualDstList)
			go b.Send(msg)
		}
	}
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
	n := b.GetNodeById(id)
	if n != nil {
		//id已经存在
		if n.IsAlive {
			//该节点存活，使存活节点下线
			b.Unregister(id)
		}
	}
	token := uuid.NewString()
	//重新给节点分配token
	n = node.NewNode(id, addr, state)
	n.RealtimeInfo = state
	//存入全局节点表
	b.NodeMap[id] = n
	log_utils.Info.Println("登录id", id)
	log_utils.Info.Printf("%v", b.NodeMap[id].Auth.Token)
	//启动保活协程
	go b.keepAlive(id)
	ctx := context.Background()
	//刷新redis存储的最新状态
	redis_utils.GlobalRedisClient.Set(ctx, n.ID, state, 60*time.Second)
	return token
}

// AuthenticateNode 节点验证
func (b *Broker) AuthenticateNode(id, addr, token string, state []byte) bool {
	b.rwm.Lock()
	defer b.rwm.Unlock()

	n := b.NodeMap[id]
	if n == nil {
		//节点未在线使该id节点上线
		log_utils.Info.Println("创建节点：", id)
		n = node.NewNode(id, addr, state)
		b.NodeMap[id] = n
		go b.keepAlive(id)
		return false
	}

	//刷新节点在线状态
	n.AliveChan <- true
	ctx := context.Background()
	//刷新redis存储的最新状态
	redis_utils.GlobalRedisClient.Set(ctx, n.ID, state, 60*time.Second)

	if n.Auth.Verify() {
		//验证通过
		return true
	}

	return false
}

// Unregister 取消某个id的节点在线状态
func (b *Broker) Unregister(id string) {
	b.NodeMap[id].AliveChan <- false
}

// GetAliveNodeIDList 获取所有在线节点ID和节点详细信息
func (b *Broker) GetAliveNodeIDList() ([]string, []node.Node) {
	b.rwm.RLock()
	defer b.rwm.RUnlock()
	var onlineNodeIdList []string
	var onlineNodeList []node.Node
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
	log_utils.Info.Println("Online Node List:", idList)
	ctx := context.Background()
	messageList, err := redis_utils.GlobalRedisClient.MGet(ctx, idList...).Result()
	if err != nil {
		log_utils.Error.Println("Get Alive Node Message Error:", err)
	}
	return messageList
}

// GetMessage 根据节点id定位一则消息
func (b *Broker) GetMessage(nodeId string) *mq.Message {
	n := b.GetNodeById(nodeId)
	msg := n.MQ.Dequeue()
	return msg
}

// KeepAlive 保持在线
func (b *Broker) keepAlive(id string) {
	n := b.GetNodeById(id)
	aliveTicker := time.NewTicker(60 * time.Second)
	for {
		select {
		case alive := <-n.AliveChan:
			if alive {
				n.IsAlive = true
				continue
			} else {
				//收到下线指令，退出保活协程
				n.IsAlive = false
				return
			}
		case <-aliveTicker.C:
			log_utils.Info.Println(id, "超时退出！")
			n.IsAlive = false
			b.Unregister(id)
			return
		}
	}
}
