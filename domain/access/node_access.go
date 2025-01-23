package access

import (
	"dominant/domain/monitor"
	"dominant/infrastructure/config"
	"github.com/Yui100901/MyGo/log_utils"
	"github.com/Yui100901/MyGo/network/mqtt_utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 11
//

type NodeClient struct {
	ID         string
	MqttClient *mqtt_utils.MQTTClient
}

func NewNodeClient(id string) *NodeClient {
	if id != "" {
		config.Config.MQTT.Node.ID = id
	}
	return &NodeClient{
		ID: id,
		MqttClient: mqtt_utils.NewMQTTClient(
			config.Config.MQTT.Node,
			map[string]byte{
				"NODE": 0,
				//"SHIP2APP/+/BASIC": 0,
			},
			func(client mqtt.Client, mqttMsg mqtt.Message) {
				payload := mqttMsg.Payload()
				topic := mqttMsg.Topic()
				log_utils.Info.Printf("%s Subscriber Received message from topic: %s\n", id, mqttMsg.Topic())
				topicParts := strings.Split(topic, "/")
				switch topicParts[0] {
				case "NODE":
					log_utils.Info.Println(string(payload))
					monitor.MainMonitor.NodeIdChan <- "123456"
				}
			}),
	}
}

func (n *NodeClient) Receive() {
	n.MqttClient.SubscribeDefault()
}

func (n *NodeClient) MQTTSend(r *mqtt_utils.MQTTPublishRequest) ([]byte, error) {
	return n.MqttClient.Publish(r)
}
