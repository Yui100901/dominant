package main

import (
	mqttutils "dominant/mqttutil"
	"dominant/mqttutil/publisher"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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
	p := publisher.NewPublisher("client-id")
	for id, topic := range idMap {
		go func() {
			for {
				body := fmt.Sprintf("%s%s", time.Now().Format("2006-01-02 15:04:05"), uuid.New())
				msg := &mqttutils.MqttMessage{
					ID:        id,
					NodeId:    id,
					Telemetry: fmt.Sprintf(`{"message":"%s"}`, body),
				}
				jsonMessage, _ := json.Marshal(msg)
				//for i := 0; i < 100; i++ {
				//client := mqttutils.NewMQTTClient(id, config.GlobalMqttConnectInfo)
				p.Publish(topic, jsonMessage)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	select {}
}
