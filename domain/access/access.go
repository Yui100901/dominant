package access

import (
	"errors"
	"github.com/Yui100901/MyGo/concurrency"
	"github.com/Yui100901/MyGo/network/http_utils"
	"github.com/Yui100901/MyGo/network/mqtt_utils"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 10
//

type Accessor struct {
	ID         string
	MqttClient *mqtt_utils.MQTTClient
	HttpClient *http_utils.HTTPClient
}

func (c *Accessor) MQTTSend(r *mqtt_utils.MQTTPublishRequest) ([]byte, error) {
	return c.MqttClient.Publish(r)
}

func (c *Accessor) HTTPSend(r *http_utils.HTTPRequest) ([]byte, error) {
	return c.HttpClient.GetResponseData(r)
}

var AccessorMap *concurrency.SafeMap[string, *Accessor]

func init() {
	AccessorMap = concurrency.NewSafeMap[string, *Accessor](32)
}

func Send(r *NodeRequest) ([]byte, error) {
	accessor, ok := AccessorMap.Get(r.NodeId)
	if !ok {
		return nil, errors.New("accessor not found")
	}
	switch r.Protocol {
	case "HTTP":
		return accessor.HTTPSend(r.HTTPRequest)
	case "MQTT":
		return accessor.MQTTSend(r.MQTTRequest)
	default:
		return nil, errors.New("not supported protocol")
	}
}

type NodeRequest struct {
	NodeId      string                         //节点id
	Protocol    string                         //协议
	HTTPRequest *http_utils.HTTPRequest        //HTTP请求
	MQTTRequest *mqtt_utils.MQTTPublishRequest //MQTT的请求
}
