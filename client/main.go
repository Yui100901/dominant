package main

import (
	"fmt"
	"github.com/Yui100901/MyGo/network/mqtt_utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//
// @Author yfy2001
// @Date 2025/1/8 21 20
//

func main() {
	p := mqtt_utils.NewMQTTClient(
		mqtt_utils.MQTTConfiguration{
			ID:       "test_client",
			URL:      "tcp://42.192.69.243:11883",
			Username: "root",
			Password: "yfy20010910",
		},
		map[string]byte{}, nil)
	sendMsg := fmt.Sprintf(`{
		Time: %s,
		Data: "MQTT_CLIENT_TEST",
	}`, time.Now().Format("2006-01-02 15:04:05"))
	p.Subscribe(map[string]byte{
		"CONTROL": 0,
	}, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("CONTROL MESSAGE: %s\n", string(msg.Payload()))
		sendMsg = string(msg.Payload())
	})
	for {
		time.Sleep(5 * time.Second)
		p.Publish(mqtt_utils.NewMQTTPublishRequest("NODE", 0, false, sendMsg))
	}

}
