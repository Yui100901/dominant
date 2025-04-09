package persistence

import (
	"dominant/infrastructure/config"
	"dominant/persistence/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//
// @Author yfy2001
// @Date 2025/3/20 08 45
//

var DB *gorm.DB

func init() {
	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		panic("failed to connect database")
	}
	err = DB.AutoMigrate(
		&model.Device{},
		&model.Telemetry{},
		&model.TelemetryLatest{},
		&model.TelemetryVirtual{})
	if err != nil {
		panic("failed to AutoMigrate")
	}
}

func ConnectDatabase() (*gorm.DB, error) {
	// 尝试连接 MySQL
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(config.Config.Mysql.ToDSN()), &gorm.Config{})
	if err == nil {
		fmt.Println("Connected to MySQL!")
		return db, nil
	}
	// 如果 MySQL 连接失败，尝试连接 SQLite
	fmt.Println("Failed to connect to MySQL, switching to SQLite...")
	db, err = gorm.Open(sqlite.Open(config.Config.Sqlite.Path), &gorm.Config{})
	if err == nil {
		fmt.Println("Connected to SQLite!")
		return db, nil
	}

	// 如果两者都失败，返回错误
	return nil, err
}
