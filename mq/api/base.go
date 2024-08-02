package api

import "dominant/mq/node"

//
// @Author yfy2001
// @Date 2024/8/1 12 41
//

var b *node.Broker

func init() {
	b = node.NewBroker()
	go b.Message()
}
