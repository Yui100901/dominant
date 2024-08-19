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
	GlobalMqttConnectInfo.MqttUrl = "tcp://42.192.69.243:11883"
	GlobalMqttConnectInfo.Username = "root"
	GlobalMqttConnectInfo.Password = "yfy20010910"
}
