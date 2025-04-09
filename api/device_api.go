package api

import (
	"dominant/domain/impl/data"
	"dominant/domain/impl/data/common"
	"dominant/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Author yfy2001
// @Date 2025/3/14 14 21
//

const devicePrefix = "/device"

func listDevice(ctx *gin.Context) {
	var deviceList []*common.Device
	global.DeviceMap.ForEach(func(id string, device *common.Device) bool {
		deviceList = append(deviceList, device)
		return true
	})
	ctx.JSON(http.StatusOK, deviceList)
}

func getDeviceById(ctx *gin.Context) {
	id := ctx.Query("id")
	global.DeviceMap.Get(id)
	ctx.JSON(http.StatusOK, id)
}

func createOrUpdateDevice(ctx *gin.Context) {

	type DeviceCreateOrUpdateCommand struct {
		ID         string `json:"id"`         //id
		Name       string `json:"name"`       //名称
		DeviceType string `json:"deviceType"` //设备类型
		EnvType    string `json:"envType"`    //环境类型
		Model      string `json:"model"`      //设备型号
	}

	var cmd DeviceCreateOrUpdateCommand
	if err := ctx.ShouldBind(cmd); err != nil {
		ctx.JSON(http.StatusOK, Error("Request Error!"))
		return
	}
	device := &common.Device{
		ID:         cmd.ID,
		Name:       cmd.Name,
		DeviceType: cmd.DeviceType,
		EnvType:    cmd.EnvType,
		Model:      cmd.Model,
	}

	if device.ID == "" {
		ctx.JSON(http.StatusOK, Error("No ID!"))
		return
	}
	global.DeviceMap.Set(device.ID, device)
	err := data.SaveDevice(device)
	if err != nil {
		ctx.JSON(http.StatusOK, Error("Create Device Failed!"))
		return
	}
}

func deleteDevice(ctx *gin.Context) {
	var idList []string
	if err := ctx.ShouldBind(idList); err != nil {
		ctx.JSON(http.StatusOK, Error("Request Error!"))
		return
	}
	for _, id := range idList {
		global.DeviceMap.Delete(id)
	}
	data.DeleteDevice(idList)
}

func AddDeviceApi() {
	WebEngine.GET(UrlPrefix+devicePrefix+"/list", listDevice)
	WebEngine.POST(UrlPrefix+devicePrefix+"/getDeviceById", getDeviceById)
	WebEngine.POST(UrlPrefix+devicePrefix+"/createOrUpdate", createOrUpdateDevice)
	WebEngine.POST(UrlPrefix+devicePrefix+"/delete", deleteDevice)
}
