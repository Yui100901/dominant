package api

import (
	"dominant/domain/access"
	"github.com/Yui100901/MyGo/network/mqtt_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Author yfy2001
// @Date 2025/1/21 15 35
//

type ShipControlCommand struct {
	ShipId  string `json:"shipId"`
	Message string `json:"message"`
}

func ShipControl(ctx *gin.Context) {
	var cmd ShipControlCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		ctx.JSON(http.StatusOK, Error("Request Error!"))
		return
	}
	if _, err := access.Send(&access.NodeRequest{
		NodeId:      cmd.ShipId,
		Protocol:    "MQTT",
		HTTPRequest: nil,
		MQTTRequest: mqtt_utils.NewMQTTPublishRequest("SHIP2APP"+cmd.ShipId, 1, false, cmd.Message),
	}); err != nil {
		ctx.JSON(http.StatusOK, Error("Send Message Error!"))
		return
	}
	ctx.JSON(http.StatusOK, Success("Send Message Success!"))
}

func AddShipApi() {
	WebEngine.POST("/api/v3/ship/control", ShipControl)
}
