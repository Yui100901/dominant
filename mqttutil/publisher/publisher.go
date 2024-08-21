package publisher

import (
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
	client := mqtt.NewClient(nil)
	return &Publisher{
		clientId: id,
		Client:   client,
	}
}

func (p *Publisher) Publish(topic string, payload any) {
	//if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
	//	panic(conn.Error())
	//}
	res := p.Client.Publish(topic, 0, false, payload)
	res.Wait()
}

func DefaultPublishHandler(client mqtt.Client, msg mqtt.Message) {

}
