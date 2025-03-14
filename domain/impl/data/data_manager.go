package data

import (
	"dominant/domain/impl/data/common"
	"dominant/global"
)

//
// @Author yfy2001
// @Date 2025/3/12 20 49
//

func SaveDeviceTelemetry(id string, m common.DeviceMessage) {
	device := m.ConvertToDevice(id)
	telemetry := m.ConvertToTelemetry(id)
	global.DeviceMap.Set(device.ID, device)
	global.TelemetryLatestMap.Set(device.ID, telemetry)
	global.TelemetryMap.Set(telemetry.ID, telemetry)
	deviceTelemetryData := common.NewDeviceTelemetryData(device, telemetry)
	global.DeviceTelemetryDataLatestMap.Set(device.ID, deviceTelemetryData)
}
