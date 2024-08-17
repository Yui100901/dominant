package subscriber

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	broker "go-iot/borker"
	"go-iot/message"
)

//
// @Author yfy2001
// @Date 2024/8/15 21 42
//

var GlobalBroker *broker.Broker
var s *Subscriber

var callback mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()
	topic := msg.Topic()
	fmt.Printf("Subscriber Received message from topic: %s\n", msg.Topic())
	var shipMessage message.ShipMessage
	json.Unmarshal(payload, &shipMessage)
	GlobalBroker.Register(shipMessage.ShipId, "device", topic, payload)
}

func init() {
	GlobalBroker = broker.NewBroker()
	s = NewSubscriber("mqtt_subscriber",
		map[string]byte{
			"SHIP2APP/+/BASIC": 0,
		},
		callback)
	go s.Subscribe()
}
