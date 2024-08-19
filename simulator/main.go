package main

import (
	"dominant/config"
	mqttutils "dominant/mqttutil"
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
		idMap[strconv.Itoa(i)] = fmt.Sprintf(`TEST/%d`, i)
	}
	for id, topic := range idMap {
		go func() {
			for {
				body := time.Now().Format("2006-01-02 15:04:05")
				msg := &mqttutils.MqttMessage{
					ID:        id,
					NodeId:    id,
					Telemetry: fmt.Sprintf(`{"message":"%s"}`, body),
				}
				jsonMessage, _ := json.Marshal(msg)
				//for i := 0; i < 100; i++ {
				client := mqttutils.NewMQTTClient(id, config.GlobalMqttConnectInfo)
				publisher.PublishTelemetry(client, topic, jsonMessage)
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}
	select {}
}
