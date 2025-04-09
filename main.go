package main

import (
	"dominant/api"
	_ "dominant/domain/access"
	"dominant/domain/impl/data"
	"dominant/infrastructure/config"
	"dominant/task"
	"fmt"
	"github.com/Yui100901/MyGo/log_utils"
)

func main() {
	//缓存所有设备数据
	data.CacheAllDevice()
	//添加接口
	api.AddTelemetryApi()
	api.AddShipApi()
	api.AddDogControlApi()
	api.AddDogDeviceApi()
	api.AddDeviceApi()
	//定时任务
	task.Task()
	err := api.WebEngine.Run(fmt.Sprintf(":%s", config.Config.App.Port))
	if err != nil {
		log_utils.Error.Fatal(err.Error())
		return
	}
}
