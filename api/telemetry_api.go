package api

import (
	"dominant/domain/impl/data/common"
	"dominant/domain/monitor"
	"dominant/global"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//
// @Author yfy2001
// @Date 2025/1/8 21 06
//

const telemetryPrefix = "/telemetry"

func getTelemetryDataInfo() []*common.DeviceTelemetryData {
	onlineNodeIds, offlineNodeIds := monitor.GetNodesList()
	var onlineNodeDataList []*common.DeviceTelemetryData
	for _, id := range onlineNodeIds {
		data, _ := global.DeviceTelemetryDataLatestMap.Get(id)
		data.IsOnline = 1
		onlineNodeDataList = append(onlineNodeDataList, data)
	}
	var offlineNodeDataList []*common.DeviceTelemetryData
	for _, id := range offlineNodeIds {
		data, _ := global.DeviceTelemetryDataLatestMap.Get(id)
		data.IsOnline = 0
		offlineNodeDataList = append(offlineNodeDataList, data)
	}
	return append(onlineNodeDataList, offlineNodeDataList...)
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
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data := getTelemetryDataInfo()
			jsonData, err := json.Marshal(data)
			_, err = fmt.Fprintf(c.Writer, "event:message\ndata:%s\n\n", jsonData)
			if err != nil {
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

func AddTelemetryApi() {
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataInfo", telemetryDataInfo)
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataInfoSlow", telemetryDataInfoSlow)
	WebEngine.GET(UrlPrefix+telemetryPrefix+"/dataEventsSlow", dataEventsSlow)
}
