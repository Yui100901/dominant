package main

import (
	"dominant/api"
	_ "dominant/domain/access"
	"dominant/infrastructure/config"
	"fmt"
	"github.com/Yui100901/MyGo/log_utils"
)

func main() {
	api.AddTelemetryApi()
	api.AddShipApi()
	api.AddDogControlApi()
	api.AddDogDeviceApi()
	err := api.WebEngine.Run(fmt.Sprintf(":%s", config.Config.App.Port))
	if err != nil {
		log_utils.Error.Fatal(err.Error())
		return
	}
}
