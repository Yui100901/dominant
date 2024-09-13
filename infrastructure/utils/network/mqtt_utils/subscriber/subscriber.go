package subscriber

import (
	"dominant/infrastructure/config"
	"dominant/infrastructure/utils/network/mqtt_utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

//
// @Author yfy2001
// @Date 2024/8/15 15 59
//

type Subscriber struct {
	clientId string
	client   mqtt.Client
	topicMap map[string]byte
	callback mqtt.MessageHandler
}

func NewSubscriber(id string, topicMap map[string]byte, callback mqtt.MessageHandler) *Subscriber {
	s := &Subscriber{
		clientId: id,
		topicMap: topicMap,
		callback: callback,
	}
	opts := mqtt.NewClientOptions()
	opts.SetClientID(id)
	//设置断开连接时自动重新连接
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(s.OnConnectHandler)
	opts.SetConnectionLostHandler(ConnectionLostHandler)
	s.client = mqtt_utils.NewMQTTClient(s.clientId, config.Config, opts)
	if conn := s.client.Connect(); conn.Wait() && conn.Error() != nil {
		log.Println(conn.Error())
		return nil
	}
	return s
}

func (s *Subscriber) Subscribe() {
	if conn := s.client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}
	res := s.client.SubscribeMultiple(s.topicMap, s.callback)
	res.Wait()
}

func (s *Subscriber) OnConnectHandler(client mqtt.Client) {
	log.Println("Subscriber Connected!")
	go client.SubscribeMultiple(s.topicMap, s.callback)
}

func ConnectionLostHandler(client mqtt.Client, err error) {
	log.Println("Subscriber Connection Lost!")
	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		log.Println(conn.Error())
	}
}

func (s *Subscriber) Disconnect() {
	s.client.Disconnect(1000)
}
