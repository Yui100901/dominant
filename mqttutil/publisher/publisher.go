package publisher

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
// @Author yfy2001
// @Date 2024/8/15 15 58
//

func PublishTelemetry(client mqtt.Client, topic string, payload any) {
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}
	res := client.Publish(topic, 0, false, payload)
	res.Wait()
}
