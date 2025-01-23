package access

import (
	"errors"
	"github.com/Yui100901/MyGo/concurrent"
	"github.com/Yui100901/MyGo/network/http_utils"
	"github.com/Yui100901/MyGo/network/mqtt_utils"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 10
//

type Receiver interface {
	Receive()
}

type Sender interface {
	HTTPSender
	MQTTSender
}

type HTTPSender interface {
	HTTPSend(r *http_utils.HTTPRequest) ([]byte, error)
}

type MQTTSender interface {
	MQTTSend(r *mqtt_utils.MQTTPublishRequest) ([]byte, error)
}

var SenderMap *concurrent.SafeMap[string, Sender]

func init() {
	SenderMap = concurrent.NewSafeMap[string, Sender](32)
}

func Send(r *NodeRequest) ([]byte, error) {
	sender, _ := SenderMap.Get(r.NodeId)
	switch r.Protocol {
	case "HTTP":
		return sender.HTTPSend(r.HTTPRequest)
	case "MQTT":
		return sender.MQTTSend(r.MQTTRequest)
	default:
		return nil, errors.New("not supported protocol")
	}
}

type NodeRequest struct {
	NodeId      string //节点id
	Protocol    string //协议
	HTTPRequest *http_utils.HTTPRequest
	MQTTRequest *mqtt_utils.MQTTPublishRequest //MQTT的请求
}
