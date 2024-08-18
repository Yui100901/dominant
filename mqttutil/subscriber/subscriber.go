package subscriber

import (
	"dominant/config"
	"dominant/mqttutil"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
// @Author yfy2001
// @Date 2024/8/15 15 59
//

type Subscriber struct {
	client   mqtt.Client
	topicMap map[string]byte
	callback mqtt.MessageHandler
}

func NewSubscriber(clientID string, topicMap map[string]byte, callback mqtt.MessageHandler) *Subscriber {
	return &Subscriber{
		client:   mqttutil.NewMQTTClient(clientID, config.GlobalMqttConnectInfo),
		topicMap: topicMap,
		callback: callback,
	}
}

func (s *Subscriber) Subscribe() {
	if conn := s.client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}
	res := s.client.SubscribeMultiple(s.topicMap, s.callback)
	res.Wait()
}
