package api

import (
	"dominant/broker"
	"dominant/mq/message"
	mqtt_utils "dominant/mqttutil"
	"dominant/mqttutil/subscriber"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

//
// @Author yfy2001
// @Date 2024/8/17 15 43
//

var s *subscriber.Subscriber

var callback mqtt.MessageHandler = func(client mqtt.Client, mqttMsg mqtt.Message) {
	payload := mqttMsg.Payload()
	topic := mqttMsg.Topic()
	log.Printf("Subscriber Received message from topic: %s\n", mqttMsg.Topic())
	var mqttMessage *mqtt_utils.MqttMessage
	json.Unmarshal(payload, mqttMessage)
	msg := message.NewMessage(mqttMessage.ID, []string{topic}, mqttMessage)
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
