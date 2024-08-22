package main

import (
	"dominant/mqttutil/publisher"
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
		idMap[strconv.Itoa(i)] = fmt.Sprintf(`SHIP2APP/%s/BASIC`, uuid.NewString())
	}
	p := publisher.NewPublisher("client-id")
	for id, topic := range idMap {
		go func() {
			for {
				//body := fmt.Sprintf("%s%s", time.Now().Format("2006-01-02 15:04:05"), uuid.New())
				//msg := mqttutils.NewMQTTMessage(id, body)
				//jsonMessage, _ := json.Marshal(msg)
				//for i := 0; i < 100; i++ {
				//client := mqttutils.NewMQTTClient(id, config.GlobalMqttConnectInfo)
				p.Publish(topic, fmt.Sprintf(`
{
		"lat":30.55946319,
			"lng":114.32081188,
			"pd_percent":57,
			"speed":0.88,
			"network_status":2,
			"network_strength": "5",
			"linux_state": 0,
			"control":1,
			"ship_id": "%s",
			"route_id" : "",
			"wq":{
			"doxygen":12.6243563423124567,
				"tur":1.762134230782336,
				"ct":0.400247645107641,
				"ph":8.893197354673101,
				"temper":12.412349645612345,
				"nh3":0.253458763197238,
				"nh4":0.201851348623502
		},
		"alert_info": {
			"image": "",
				"time": 1709866669.903657
		}
}`, id))
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}
	select {}
}
