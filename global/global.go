package global

import (
	"dominant/domain/impl/data/common"
	"github.com/Yui100901/MyGo/concurrent"
)

//
// @Author yfy2001
// @Date 2025/1/8 20 43
//

var DeviceMap *concurrent.SafeMap[string, *common.Device]
var TelemetryLatestMap *concurrent.SafeMap[string, *common.Telemetry]
var TelemetryMap *concurrent.SafeMap[string, *common.Telemetry]
var DeviceTelemetryDataLatestMap *concurrent.SafeMap[string, *common.DeviceTelemetryData]
var VirtualDeviceTelemetryDataMap *concurrent.SafeMap[string, *common.DeviceTelemetryData]

func init() {
	DeviceMap = concurrent.NewSafeMap[string, *common.Device](32)
	TelemetryLatestMap = concurrent.NewSafeMap[string, *common.Telemetry](32)
	TelemetryMap = concurrent.NewSafeMap[string, *common.Telemetry](32)
	DeviceTelemetryDataLatestMap = concurrent.NewSafeMap[string, *common.DeviceTelemetryData](32)
	VirtualDeviceTelemetryDataMap = concurrent.NewSafeMap[string, *common.DeviceTelemetryData](32)
}
