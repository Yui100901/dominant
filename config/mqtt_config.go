package config

//
// @Author yfy2001
// @Date 2024/8/17 13 43
//

var GlobalMqttConnectInfo MqttConnectInfo

type MqttConnectInfo struct {
	MqttUrl  string
	Username string
	Password string
}

func init() {
	GlobalMqttConnectInfo.MqttUrl = "tcp://192.168.1.200:11883"
	GlobalMqttConnectInfo.Username = "root"
	GlobalMqttConnectInfo.Password = "123456"
}
