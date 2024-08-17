package broker

//
// @Author yfy2001
// @Date 2024/8/1 12 41
//

// GlobalBroker 全局broker负责接收分发消息
var GlobalBroker *Broker

func init() {
	GlobalBroker = NewBroker()
	go GlobalBroker.Distribute()
}
