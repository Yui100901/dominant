package model

import (
	"time"
)

//
// @Author yfy2001
// @Date 2025/4/3 13 12
//

// TelemetryLatest 最新表以deviceId作为主键
// 只保留设备的最新数据
type TelemetryLatest struct {
	ID            string    `json:"id"`                         //id
	DeviceID      string    `json:"deviceId" gorm:"primaryKey"` //设备id
	TelemetryTime time.Time `json:"telemetryTime"`              //遥测上传时间
	Position      *Position `json:"position" gorm:"type:json"`  //位置
	Status        *Status   `json:"status" gorm:"type:json"`    //状态
	RawData       *JsonObj  `json:"rawData" gorm:"type:json"`   //原始数据
}

func (TelemetryLatest) TableName() string {
	return "telemetry_latest"
}
