package publisher

import (
	"dominant/infrastructure/config"
	"dominant/infrastructure/utils/log_utils"
	"dominant/infrastructure/utils/network/mqtt_utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
// @Author yfy2001
// @Date 2024/8/15 15 58
//

type Publisher struct {
	clientId string
	Client   mqtt.Client
}

func NewPublisher(id string) *Publisher {
	opts := mqtt.NewClientOptions()
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(DefaultPublishHandler)
	opts.SetOnConnectHandler(OnConnectHandler)
	opts.SetConnectionLostHandler(ConnectionLostHandler)
	client := mqtt_utils.NewMQTTClient(id, config.Config, opts)
	return &Publisher{
		clientId: id,
		Client:   client,
	}
}

func (p *Publisher) Publish(topic string, payload any) {
	if conn := p.Client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}
	res := p.Client.Publish(topic, 0, false, payload)
	res.Wait()
}

func OnConnectHandler(client mqtt.Client) {
	log_utils.Info.Println("Publisher Connected!")
}

func ConnectionLostHandler(client mqtt.Client, err error) {
	log_utils.Info.Println("Publisher Connection Lost!")
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		log_utils.Error.Println(conn.Error())
	}
}

func DefaultPublishHandler(client mqtt.Client, msg mqtt.Message) {

}
