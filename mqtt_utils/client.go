package mqtt_utils

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go-iot/base"
)

//
// @Author yfy2001
// @Date 2024/8/15 16 00
//

func NewMQTTClient(clientID string, info base.MqttConnectInfo) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.SetClientID(clientID)
	opts.AddBroker(info.MqttUrl)
	opts.SetUsername(info.Username)
	opts.SetPassword(info.Password)
	opts.SetDefaultPublishHandler(defaultPublishHandler)
	opts.SetOnConnectHandler(onConnectHandler)
	opts.SetConnectionLostHandler(connectionLostHandler)
	client := mqtt.NewClient(opts)
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}
	return client
}