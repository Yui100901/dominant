package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

//
// @Author yfy2001
// @Date 2025/3/31 11 01
//

type Model interface {
	TableName() string // 确保所有泛型类型都有表名
}

type JsonObj struct {
	Data any `json:"data"`
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的序列化
func (p *JsonObj) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取时的反序列化
func (p *JsonObj) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("类型转换失败: %v", value)
	}
	return json.Unmarshal(b, p)
}
