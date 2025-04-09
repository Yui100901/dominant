package model

import (
	"time"
)

//
// @Author yfy2001
// @Date 2025/4/3 13 12
//

// TelemetryVirtual 虚拟表id为主键，允许出现同一个设备在不同状态的历史
type TelemetryVirtual struct {
	ID            string    `json:"id" gorm:"primaryKey"`      //id
	DeviceID      string    `json:"deviceId"`                  //设备id
	TelemetryTime time.Time `json:"telemetryTime"`             //遥测上传时间
	Position      *Position `json:"position" gorm:"type:json"` //位置
	Status        *Status   `json:"status" gorm:"type:json"`   //状态
	RawData       *JsonObj  `json:"rawData" gorm:"type:json"`  //原始数据
}

func (TelemetryVirtual) TableName() string {
	return "telemetry_virtual"
}
