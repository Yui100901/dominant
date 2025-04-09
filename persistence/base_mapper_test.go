package persistence

import (
	"dominant/persistence/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

//
// @Author yfy2001
// @Date 2025/3/31 11 03
//

// 初始化内存数据库
func setupDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("数据库初始化失败: %v", err)
	}
	if err := db.AutoMigrate(&model.Device{}); err != nil {
		t.Fatalf("数据库迁移失败: %v", err)
	}
	return db
}

// 打印设备表内容
func printDeviceTable(db *gorm.DB, t *testing.T) {
	var devices []model.Device
	result := db.Find(&devices)
	if result.Error != nil {
		t.Logf("查询失败: %v", result.Error)
		return
	}

	t.Log("\n当前设备表:")
	if len(devices) == 0 {
		t.Log("(空)")
		return
	}

	fmt.Printf("%-4s | %-15s | %-10s | %-8s | %-8s\n",
		"ID", "Name", "DeviceType", "EnvType", "Model")
	fmt.Println("--------------------------------------------------------------")
	for _, d := range devices {
		fmt.Printf("%-4s | %-15s | %-10s | %-8s | %-8s\n",
			d.ID, d.Name, d.DeviceType, d.EnvType, d.Model)
	}
	fmt.Println()
}

func TestSaveOrUpdate(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestSaveOrUpdate ===")

	// 测试插入
	device := &model.Device{ID: "1", Name: "Device1"}
	err := SaveOrUpdate(device)
	if err != nil {
		t.Fatalf("保存失败: %v", err)
	}
	t.Log("插入后:")
	printDeviceTable(db, t)

	// 测试更新
	device.Name = "UpdatedDevice1"
	err = SaveOrUpdate(device)
	if err != nil {
		t.Fatalf("更新失败: %v", err)
	}
	t.Log("更新后:")
	printDeviceTable(db, t)
}

func TestBatchInsert(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestBatchInsert ===")

	devices := []model.Device{
		{ID: "1", Name: "Device1"},
		{ID: "2", Name: "Device2"},
	}

	t.Log("插入前:")
	printDeviceTable(db, t)

	err := BatchInsert(devices)
	if err != nil {
		t.Fatalf("批量插入失败: %v", err)
	}
	t.Log("插入后:")
	printDeviceTable(db, t)
}

func TestUpdate(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestUpdate ===")

	db.Create(&model.Device{ID: "1", Name: "OldName", DeviceType: "TypeA"})
	t.Log("初始数据:")
	printDeviceTable(db, t)

	updates := map[string]interface{}{"name": "NewName"}
	err := Update[model.Device]("1", updates)
	if err != nil {
		t.Fatalf("更新失败: %v", err)
	}
	t.Log("更新后:")
	printDeviceTable(db, t)
}

func TestGetList(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestGetList ===")

	db.Create(&model.Device{ID: "1"})
	db.Create(&model.Device{ID: "2"})

	list, err := GetList[model.Device]()
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	t.Logf("获取到 %d 条记录:", len(list))
	printDeviceTable(db, t)
}

func TestGetByID(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestGetByID ===")

	db.Create(&model.Device{ID: "1", Name: "TargetDevice"})

	// 存在的情况
	result, err := GetByID[model.Device]("1")
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	t.Logf("查询结果: %+v", *result)

	// 不存在的情况
	_, err = GetByID[model.Device]("999")
	if err != gorm.ErrRecordNotFound {
		t.Errorf("期望得到未找到错误，实际得到: %v", err)
	}
}

func TestGetPaginatedList(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestGetPaginatedList ===")

	// 插入10条测试数据
	for i := 1; i <= 10; i++ {
		db.Create(&model.Device{ID: fmt.Sprint(i), Name: fmt.Sprintf("Device%d", i)})
	}

	t.Log("完整列表:")
	printDeviceTable(db, t)

	// 测试分页
	page1, err := GetPaginatedList[model.Device](1, 3)
	if err != nil {
		t.Fatalf("分页查询失败: %v", err)
	}
	t.Logf("第1页（3条）:")
	for _, d := range page1 {
		t.Logf("ID: %s, Name: %s", d.ID, d.Name)
	}

	page2, err := GetPaginatedList[model.Device](2, 3)
	if err != nil {
		t.Fatalf("分页查询失败: %v", err)
	}
	t.Logf("\n第2页（3条）:")
	for _, d := range page2 {
		t.Logf("ID: %s, Name: %s", d.ID, d.Name)
	}
}

func TestGetByCondition(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestGetByCondition ===")

	db.Create(&model.Device{ID: "1", DeviceType: "TypeA", EnvType: "Indoor"})
	db.Create(&model.Device{ID: "2", DeviceType: "TypeB", EnvType: "Outdoor"})
	db.Create(&model.Device{ID: "3", DeviceType: "TypeA", EnvType: "Outdoor"})

	// 多条件查询
	conditions := map[string]interface{}{
		"device_type": "TypeA",
		"env_type":    "Outdoor",
	}
	list, err := GetByCondition[model.Device](conditions)
	if err != nil {
		t.Fatalf("条件查询失败: %v", err)
	}
	t.Log("查询结果:")
	for _, d := range list {
		t.Logf("ID: %s, Type: %s, Env: %s", d.ID, d.DeviceType, d.EnvType)
	}
}

func TestGetSortedList(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestGetSortedList ===")

	db.Create(&model.Device{ID: "3", Name: "DeviceC"})
	db.Create(&model.Device{ID: "1", Name: "DeviceA"})
	db.Create(&model.Device{ID: "2", Name: "DeviceB"})

	// 按ID升序
	list, err := GetSortedList[model.Device]("id asc")
	if err != nil {
		t.Fatalf("排序查询失败: %v", err)
	}
	t.Log("按ID升序:")
	for _, d := range list {
		t.Logf("ID: %s, Name: %s", d.ID, d.Name)
	}

	// 按名称降序
	list, err = GetSortedList[model.Device]("name desc")
	if err != nil {
		t.Fatalf("排序查询失败: %v", err)
	}
	t.Log("\n按名称降序:")
	for _, d := range list {
		t.Logf("ID: %s, Name: %s", d.ID, d.Name)
	}
}

func TestDeleteByID(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestDeleteByID ===")

	db.Create(&model.Device{ID: "1"})
	db.Create(&model.Device{ID: "2"})

	t.Log("删除前:")
	printDeviceTable(db, t)

	err := DeleteByID[model.Device]("1")
	if err != nil {
		t.Fatalf("删除失败: %v", err)
	}
	t.Log("删除后:")
	printDeviceTable(db, t)
}

func TestBatchDelete(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestBatchDelete ===")

	db.Create(&model.Device{ID: "1", DeviceType: "TypeA"})
	db.Create(&model.Device{ID: "2", DeviceType: "TypeA"})
	db.Create(&model.Device{ID: "3", DeviceType: "TypeB"})

	t.Log("删除前:")
	printDeviceTable(db, t)

	err := BatchDelete[model.Device](map[string]interface{}{"device_type": "TypeA"})
	if err != nil {
		t.Fatalf("批量删除失败: %v", err)
	}
	t.Log("删除后:")
	printDeviceTable(db, t)
}

func TestExists(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestExists ===")

	// 初始不存在
	exists, err := Exists[model.Device](map[string]interface{}{"id": "1"})
	if err != nil {
		t.Fatalf("检查存在性失败: %v", err)
	}
	if exists {
		t.Error("不应存在记录")
	}

	db.Create(&model.Device{ID: "1"})
	exists, err = Exists[model.Device](map[string]interface{}{"id": "1"})
	if err != nil {
		t.Fatalf("检查存在性失败: %v", err)
	}
	if !exists {
		t.Error("应存在记录")
	}
}

func TestCount(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestCount ===")

	// 初始计数
	count, err := Count[model.Device]()
	if err != nil {
		t.Fatalf("计数失败: %v", err)
	}
	if count != 0 {
		t.Errorf("期望0条，实际%d条", count)
	}

	db.Create(&model.Device{ID: "1"})
	count, err = Count[model.Device]()
	if err != nil {
		t.Fatalf("计数失败: %v", err)
	}
	if count != 1 {
		t.Errorf("期望1条，实际%d条", count)
	}
}

func TestCountByCondition(t *testing.T) {
	db := setupDB(t)
	DB = db
	t.Log("=== TestCountByCondition ===")

	db.Create(&model.Device{ID: "1", DeviceType: "TypeA"})
	db.Create(&model.Device{ID: "2", DeviceType: "TypeA"})
	db.Create(&model.Device{ID: "3", DeviceType: "TypeB"})

	count, err := CountByCondition[model.Device](map[string]interface{}{"device_type": "TypeA"})
	if err != nil {
		t.Fatalf("条件计数失败: %v", err)
	}
	if count != 2 {
		t.Errorf("期望2条，实际%d条", count)
	}
}
