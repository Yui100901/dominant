package mqtt_utils

import (
	"dominant/infrastructure/config"
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

func NewMQTTClient(clientID string, config config.Configuration, opts *mqtt.ClientOptions) mqtt.Client {
	opts.SetClientID(clientID)
	opts.AddBroker(config.MQTT.URL)
	opts.SetUsername(config.MQTT.Username)
	opts.SetPassword(config.MQTT.Password)
	client := mqtt.NewClient(opts)
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		log.Println(conn.Error())
		return nil
	}
	return client
}
