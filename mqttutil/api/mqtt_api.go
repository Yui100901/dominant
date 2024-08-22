package api

import (
	"dominant/broker"
	"dominant/mq/message"
	mqttutils "dominant/mqttutil"
	"dominant/mqttutil/subscriber"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	mqttMessage := new(mqttutils.MQTTMessage)
	err := json.Unmarshal(payload, mqttMessage)
	if err != nil {
		log.Fatal(err)
	}
	msg := message.NewMessage(topic, "status", mqttMessage.NodeId, []string{topic}, mqttMessage)
	broker.GlobalBroker.Register(mqttMessage.NodeId, msg.Topic, payload)
}

func init() {
	s = subscriber.NewSubscriber(
		"mqtt_subscriber",
		map[string]byte{
			"TEST/+":           0,
			"SHIP2APP/+/BASIC": 0,
		},
		callback)
	//time.Sleep(3000 * time.Millisecond)
	//s.Disconnect()
}

func GetNodeStatusList(c *gin.Context) {
	msgList := broker.GlobalBroker.GetAliveNodeMessage()
	var mqttMessageList []*mqttutils.MQTTMessage
	for _, msg := range msgList {
		stringMessage := msg.(string)
		mqttMessage := new(mqttutils.MQTTMessage)
		json.Unmarshal([]byte(stringMessage), mqttMessage)
		mqttMessageList = append(mqttMessageList, mqttMessage)
	}
	c.JSON(http.StatusOK, msgList)
}
