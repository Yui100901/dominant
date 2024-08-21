package mqttutil

import (
	"dominant/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

//
// @Author yfy2001
// @Date 2024/8/15 16 00
//

type MQTTHandler interface {
	OnConnectHandler(client mqtt.Client)
	ConnectionLostHandler(client mqtt.Client, err error)
}

type MQTTClient struct {
	ClientId string
	client   mqtt.Client
}

func NewMQTTClient(clientID string, info config.MqttConnectInfo, opts *mqtt.ClientOptions) mqtt.Client {
	opts.SetClientID(clientID)
	opts.AddBroker(info.MqttUrl)
	opts.SetUsername(info.Username)
	opts.SetPassword(info.Password)
	client := mqtt.NewClient(opts)
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		log.Println(conn.Error())
		return nil
	}
	return client
}
