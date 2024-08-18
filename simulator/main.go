package main

import (
	"dominant/config"
	mqtt_utils "dominant/mqttutil"
	"dominant/mqttutil/publisher"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

//
// @Author yfy2001
// @Date 2024/8/15 09 25
//

func main() {

	idMap := make(map[string]string)
	for i := 0; i < 10; i++ {
		idMap[strconv.Itoa(i)] = fmt.Sprintf(`SHIP2APP/%d/BASIC`, i)
	}
	for id, topic := range idMap {
		msg := &mqtt_utils.MqttMessage{
			ID:        "jdaidj",
			NodeId:    "123456",
			Telemetry: `{"message":"hello world"}`,
		}
		jsonMessage, _ := json.Marshal(msg)
		go func() {
			for {
				//for i := 0; i < 100; i++ {
				client := mqtt_utils.NewMQTTClient(id, config.GlobalMqttConnectInfoBase)
				publisher.PublishTelemetry(client, topic, jsonMessage)
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}
	select {}
}