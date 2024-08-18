package api

import (
	"dominant/broker"
	"dominant/mq/message"
	"dominant/mqttutil/subscriber"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
// @Author yfy2001
// @Date 2024/8/17 15 43
//

var s *subscriber.Subscriber

var callback mqtt.MessageHandler = func(client mqtt.Client, mqttMsg mqtt.Message) {
	payload := mqttMsg.Payload()
	topic := mqttMsg.Topic()
	fmt.Printf("Subscriber Received message from topic: %s\n", mqttMsg.Topic())
	var jsonMessage *message.Message
	json.Unmarshal(payload, jsonMessage)
	msg := message.NewMessage(jsonMessage.ID, []string{topic}, jsonMessage)
	broker.GlobalBroker.MainMQ.Enqueue(msg)
	broker.GlobalBroker.Register(msg.ID, "device")
}

func init() {
	s = subscriber.NewSubscriber("mqtt_subscriber",
		map[string]byte{
			"SHIP2APP/+/BASIC": 0,
		},
		callback)
	go s.Subscribe()
}
