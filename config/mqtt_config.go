package config

//
// @Author yfy2001
// @Date 2024/8/17 13 43
//

var GlobalMqttConnectInfoBase MqttConnectInfo

type MqttConnectInfo struct {
	MqttUrl  string
	Username string
	Password string
}

func init() {
	GlobalMqttConnectInfoBase.MqttUrl = "tcp://192.168.1.200:11883"
	GlobalMqttConnectInfoBase.Username = "root"
	GlobalMqttConnectInfoBase.Password = "123456"
}
