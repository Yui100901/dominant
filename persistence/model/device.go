package model

//
// @Author yfy2001
// @Date 2025/3/19 17 27
//

type Device struct {
	ID         string `gorm:"primaryKey"` //id
	Name       string //名称
	DeviceType string //设备类型
	EnvType    string //环境类型
	Model      string //设备型号
}

func (Device) TableName() string {
	return "device"
}
