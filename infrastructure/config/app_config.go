package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//
// @Author yfy2001
// @Date 2024/9/8 13 24
//

type Configuration struct {
	App   AppConfiguration   `yaml:"app"`
	MQTT  MQTTConfiguration  `yaml:"mqtt"`
	Redis RedisConfiguration `yaml:"redis"`
}

type AppConfiguration struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type MQTTConfiguration struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfiguration struct {
	URL          string `yaml:"url"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"DB"`
	DialTimeout  int    `yaml:"DialTimeout"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
	PoolSize     int    `yaml:"PoolSize"`
	PoolTimeout  int    `yaml:"PoolTimeout"`
}

var Config Configuration

func init() {
	// 获取环境变量
	env := "dev" // 默认环境为 dev

	// 设置配置文件名和路径
	viper.SetConfigName(fmt.Sprintf("infrastructure.config-%s", env))
	viper.AddConfigPath("./infrastructure.config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	// 解析配置文件到 Configuration 结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}
