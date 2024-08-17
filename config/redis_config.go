package config

//
// @Author yfy2001
// @Date 2024/8/17 13 43
//

var GlobalRedisConnectInfoBase RedisConnectInfo

type RedisConnectInfo struct {
	RedisUrl string
	Username string
	Password string
}

func init() {

	GlobalRedisConnectInfoBase.RedisUrl = "192.168.1.200:6379"
	GlobalRedisConnectInfoBase.Username = ""
	GlobalRedisConnectInfoBase.Password = "rediszmkj123456"
}
