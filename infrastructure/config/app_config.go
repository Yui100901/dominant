package config

import (
	"github.com/Yui100901/MyGo/log_utils"
	"github.com/Yui100901/MyGo/network/mqtt_utils"

	"fmt"
	"github.com/spf13/viper"
)

//
// @Author yfy2001
// @Date 2024/9/8 13 24
//

type Configuration struct {
	App  AppConfiguration `yaml:"app"`
	MQTT struct {
		Ship mqtt_utils.MQTTConfiguration `yaml:"ship"`
		Dog  mqtt_utils.MQTTConfiguration `yaml:"dog"`
	} `yaml:"mqtt"`
	Mysql  MysqlConfiguration  `yaml:"mysql"`
	Sqlite SqliteConfiguration `yaml:"sqlite"`
	Redis  RedisConfiguration  `yaml:"redis"`
}

type AppConfiguration struct {
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type MysqlConfiguration struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func (mc *MysqlConfiguration) ToDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mc.Username,
		mc.Password,
		mc.Url,
		mc.DBName,
	)
	return dsn
}

type SqliteConfiguration struct {
	Path string `yaml:"path"`
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
	viper.SetConfigName(fmt.Sprintf("config-%s", env))
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log_utils.Error.Fatalf("无法读取配置文件: %v", err)
	}

	// 解析配置文件到 Configuration 结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log_utils.Error.Fatalf("无法解析配置文件: %v", err)
	}

}
