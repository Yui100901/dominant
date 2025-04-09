package api

import (
	"dominant/domain/impl/data"
	"dominant/domain/impl/data/common"
	"dominant/persistence/model"
	"encoding/json"
	"fmt"
	"github.com/Yui100901/MyGo/log_utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

//
// @Author yfy2001
// @Date 2025/1/8 21 06
//

const telemetryPrefix = "/telemetry"

func getTelemetryDataInfo() []*common.DeviceTelemetryData {
	var nodeDataList []*common.DeviceTelemetryData

	var telemetryList []common.Telemetry
	//获取最新遥测数据
	telemetryLatestList := data.GetTelemetryLatestList()
	//获取虚拟遥测数据
	telemetryVirtualList := data.GetTelemetryVirtualList()
	if len(telemetryLatestList) != 0 {
		telemetryList = append(telemetryList, telemetryLatestList...)
	}
	if len(telemetryVirtualList) != 0 {
		telemetryList = append(telemetryList, telemetryVirtualList...)
	}

	//处理所有遥测数据
	for _, telemetry := range telemetryList {
		handleDeviceId := func(dId string) string {
			deviceIdAliasMap := make(map[string]string)
			deviceIdAliasMap["Vehicle-5"] = "Vehicle-4"
			deviceIdAliasMap["Vehicle-6"] = "Vehicle-3"
			id, ok := deviceIdAliasMap[dId]
			if !ok {
				id = dId
			}
			return id
		}
		deviceId := handleDeviceId(telemetry.DeviceID)
		device := data.GetDeviceByIdFromMap(deviceId)
		device.ID = deviceId
		deviceTelemetryData := common.NewDeviceTelemetryData(device, &telemetry)
		nodeDataList = append(nodeDataList, deviceTelemetryData)
	}

	//log_utils.Info.Printf("nodeDataList\n%v", nodeDataList)
	return nodeDataList
}

// telemetryDataInfo 遥测数据获取HTTP接口
func telemetryDataInfo(ctx *gin.Context) {
	ctx.JSON(200, Success(getTelemetryDataInfo()))

}

// telemetryDataInfoSlow 遥测数据获取HTTP慢接口
func telemetryDataInfoSlow(ctx *gin.Context) {
	ctx.JSON(200, Success(getTelemetryDataInfo()))
}

func dataEventsSlow(c *gin.Context) {
	deviceType := c.Query("deviceType")
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	deviceTypeMap := make(map[string]string)
	deviceTypeMap["dog"] = "机器狗"
	mappedDeviceType, ok := deviceTypeMap[deviceType]
	if !ok {
		mappedDeviceType = "all"
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			telemetryData := getTelemetryDataInfo()
			filteredData := Filter(telemetryData, func(telemetryData *common.DeviceTelemetryData) bool {
				if mappedDeviceType == "all" {
					return true
				}
				return mappedDeviceType == telemetryData.DeviceType
			})
			jsonData, err := json.Marshal(filteredData)
			_, err = fmt.Fprintf(c.Writer, "event:message\ndata:%s\n\n", jsonData)
			if err != nil {
				log_utils.Error.Println("SSE write error:", err)
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

// Filter 函数，支持泛型
func Filter[T any](data []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range data {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func createOrUpdateTelemetryVirtual(ctx *gin.Context) {
	type CreateOrUpdateVirtualTelemetryCommand struct {
		ID        string  `json:"id"`        //id
		DeviceId  string  `json:"deviceId"`  //关联的设备id
		Longitude float64 `json:"longitude"` //经度
		Latitude  float64 `json:"latitude"`  //纬度
		Height    float64 `json:"height"`    //高度
		Velocity  float64 `json:"velocity"`  //速度
	}
	var cmd CreateOrUpdateVirtualTelemetryCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		log_utils.Info.Println(err.Error())
		ctx.JSON(http.StatusOK, Error("Request Error!"))
		return
	}
	var id string
	if cmd.ID != "" {
		id = cmd.ID
	} else {
		id = uuid.NewString()
	}
	telemetry := &common.Telemetry{
		ID:            id,
		DeviceID:      cmd.DeviceId,
		TelemetryTime: time.Now(),
		Position: &common.Position{
			Longitude: cmd.Longitude,
			Latitude:  cmd.Latitude,
			Height:    cmd.Height,
		},
		Status: &common.Status{
			Velocity: cmd.Velocity,
		},
		RawData: &model.JsonObj{},
	}
	data.SaveTelemetryVirtual(telemetry)
}

func deleteVirtualTelemetry(ctx *gin.Context) {

}

func AddTelemetryApi() {
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataInfo", telemetryDataInfo)
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataInfoSlow", telemetryDataInfoSlow)
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataEventsSlow", dataEventsSlow)
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/createOrUpdateTelemetryVirtual", createOrUpdateTelemetryVirtual)
}
