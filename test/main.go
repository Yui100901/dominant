package main

import (
	"dominant/authentication"
	"dominant/config"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.NewString())
	c := code.NewAuthentication(100)
	fmt.Println(c)
	// 使用配置
	fmt.Printf("应用名称: %s\n", config.Config.App.Name)
	fmt.Printf("应用版本: %s\n", config.Config.App.Version)
	fmt.Printf("MQTT Config: %+v\n", config.Config.MQTT)
	fmt.Printf("Redis Config: %+v\n", config.Config.Redis)
}
