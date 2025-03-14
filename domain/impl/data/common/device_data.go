package common

import (
	"github.com/google/uuid"
)

//
// @Author yfy2001
// @Date 2025/1/21 14 10
//

type DeviceTelemetryData struct {
	Id            string  `json:"id"`
	DeviceId      string  `json:"deviceId"`
	DeviceType    string  `json:"deviceType"`
	Name          string  `json:"name"`
	Model         string  `json:"model"`
	IsOnline      int     `json:"isOnline"`
	TelemetryTime string  `json:"telemetryTime"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	Height        float64 `json:"height"`
	Velocity      float64 `json:"velocity"`
	Battery       int     `json:"battery"`
	BatteryLife   int     `json:"batteryLife"`
	EnvType       string  `json:"envType"`
	VideoLink     string  `json:"videoLink"`
	IsVideoStream bool    `json:"isVideoStream"`
	RawData       any     `json:"rawData"`
}

func NewDeviceTelemetryData(device *Device, telemetry *Telemetry) *DeviceTelemetryData {
	dataId := uuid.NewString()
	return &DeviceTelemetryData{
		Id:            dataId,
		DeviceId:      device.ID,
		DeviceType:    device.DeviceType,
		EnvType:       device.EnvType,
		Name:          device.Name,
		Model:         device.Model,
		IsOnline:      1,
		TelemetryTime: telemetry.TelemetryTime.Format("2006-01-02 15:04:05"),
		Longitude:     telemetry.Position.Longitude,
		Latitude:      telemetry.Position.Latitude,
		Height:        telemetry.Position.Height,
		Velocity:      telemetry.Status.Velocity,
		Battery:       telemetry.Status.Battery,
		BatteryLife:   telemetry.Status.BatteryLife,
		VideoLink:     deviceVideoLinkMap[device.ID],
		IsVideoStream: false,
		RawData:       telemetry.RawData,
	}
}
