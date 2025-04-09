package common

import (
	"dominant/persistence/model"
	"time"
)

//
// @Author yfy2001
// @Date 2025/1/8 10 59
//

var deviceVideoLinkMap map[string]string

type DeviceMessage interface {
	ConvertToTelemetry(deviceId string) *Telemetry
}

type Device struct {
	ID         string `json:"id"`         //id
	Name       string `json:"name"`       //名称
	DeviceType string `json:"deviceType"` //设备类型
	EnvType    string `json:"envType"`    //环境类型
	Model      string `json:"model"`      //设备型号
}

type Telemetry struct {
	ID            string         `json:"id"`            //id
	DeviceID      string         `json:"deviceId"`      //设备id
	TelemetryTime time.Time      `json:"telemetryTime"` //遥测上传时间
	Position      *Position      `json:"position"`      //位置
	Status        *Status        `json:"status"`        //状态
	RawData       *model.JsonObj `json:"rawData"`       //原始数据
}

type Position struct {
	Longitude float64 `json:"longitude"` //经度
	Latitude  float64 `json:"latitude"`  //纬度
	Height    float64 `json:"height"`    //高度
}

type Status struct {
	Velocity    float64 `json:"velocity"`    //速度
	Battery     int     `json:"battery"`     //电量
	BatteryLife int     `json:"batteryLife"` //续航时间
}

func init() {
	deviceVideoLinkMap = make(map[string]string)
	deviceVideoLinkMap["9dd4299f15f74ffb8d52a39ffc1d2dc7"] = "http://liveplay2.orca-tech.cn/live/r91_0.flv"
}
