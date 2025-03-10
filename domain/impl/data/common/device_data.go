package common

//
// @Author yfy2001
// @Date 2025/1/21 14 10
//

type DeviceData struct {
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
