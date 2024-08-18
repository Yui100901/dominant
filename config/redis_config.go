package config

//
// @Author yfy2001
// @Date 2024/8/17 13 43
//

var GlobalRedisConnectInfo RedisConnectInfo

type RedisConnectInfo struct {
	RedisUrl string
	Username string
	Password string
}

func init() {

	GlobalRedisConnectInfo.RedisUrl = "192.168.1.200:6379"
	GlobalRedisConnectInfo.Username = ""
	GlobalRedisConnectInfo.Password = "123456"
}
