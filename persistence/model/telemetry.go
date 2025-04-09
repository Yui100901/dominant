package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

//
// @Author yfy2001
// @Date 2025/3/31 10 34
//

// Telemetry 历史表，id为主键存储历史数据
type Telemetry struct {
	ID            string    `json:"id" gorm:"primaryKey"`      //id
	DeviceID      string    `json:"deviceId"`                  //设备id
	TelemetryTime time.Time `json:"telemetryTime"`             //遥测上传时间
	Position      *Position `json:"position" gorm:"type:json"` //位置
	Status        *Status   `json:"status" gorm:"type:json"`   //状态
	RawData       *JsonObj  `json:"rawData" gorm:"type:json"`  //原始数据
}

func (Telemetry) TableName() string {
	return "telemetry"
}

type Position struct {
	Longitude float64 `json:"longitude"` //经度
	Latitude  float64 `json:"latitude"`  //纬度
	Height    float64 `json:"height"`    //高度
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的序列化
func (p *Position) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取时的反序列化
func (p *Position) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("类型转换失败: %v", value)
	}
	return json.Unmarshal(b, p)
}

type Status struct {
	Velocity    float64 `json:"velocity"`    //速度
	Battery     int     `json:"battery"`     //电量
	BatteryLife int     `json:"batteryLife"` //续航时间
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的序列化
func (p *Status) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取时的反序列化
func (p *Status) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("类型转换失败: %v", value)
	}
	return json.Unmarshal(b, p)
}
