package api

import (
	"dominant/broker"
	"dominant/mq/message"
	"dominant/mqtt_utils/subscriber"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
// @Author yfy2001
// @Date 2024/8/17 15 43
//

var s *subscriber.Subscriber

var callback mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()
	//topic := msg.Topic()
	fmt.Printf("Subscriber Received message from topic: %s\n", msg.Topic())
	var jsonMessage *message.Message
	json.Unmarshal(payload, jsonMessage)
	broker.GlobalBroker.MainMQ.Enqueue(jsonMessage)
	broker.GlobalBroker.Register(jsonMessage.ID, "device")
}

func init() {
	s = subscriber.NewSubscriber("mqtt_subscriber",
		map[string]byte{
			"SHIP2APP/+/BASIC": 0,
		},
		callback)
	go s.Subscribe()
}
