package persistence

import (
	"dominant/persistence/model"
	"fmt"
)

//
// @Author yfy2001
// @Date 2025/3/31 10 56
//

// SaveOrUpdate 插入或更新
func SaveOrUpdate[T model.Model](t *T) error {
	result := DB.Save(t)
	return result.Error
}

// BatchInsert 批量插入
func BatchInsert[T model.Model](records []T) error {
	result := DB.Create(&records)
	return result.Error
}

// Update 更新部分字段
func Update[T model.Model](id string, updates map[string]interface{}) error {
	var t T
	result := DB.Model(&t).Where("id = ?", id).Updates(updates)
	return result.Error
}

// GetList 批量查询
func GetList[T model.Model]() ([]T, error) {
	var list []T
	result := DB.Find(&list)
	return list, result.Error
}

// GetByID 根据id查询
func GetByID[T model.Model](id string) (*T, error) {
	var t T
	result := DB.First(&t, "id = ?", id)
	return &t, result.Error
}

// GetPaginatedList 分页查询
func GetPaginatedList[T model.Model](page, pageSize int) ([]T, error) {
	var list []T
	offset := (page - 1) * pageSize
	result := DB.Limit(pageSize).Offset(offset).Find(&list)
	return list, result.Error
}

// GetByCondition 条件查询
func GetByCondition[T model.Model](conditions map[string]interface{}) ([]T, error) {
	var list []T
	result := DB.Where(conditions).Find(&list)
	return list, result.Error
}

// GetSortedList 查询并排序
func GetSortedList[T model.Model](orderBy string) ([]T, error) {
	var list []T
	result := DB.Order(orderBy).Find(&list)
	return list, result.Error
}

// DeleteByID 根据id删除
func DeleteByID[T model.Model](id string) error {
	var t T // 创建类型实例
	result := DB.Delete(&t, "id = ?", id)
	return result.Error
}

// BatchDelete 批量删除
func BatchDelete[T model.Model](conditions map[string]interface{}) error {
	var t T
	result := DB.Where(conditions).Delete(&t)
	return result.Error
}

// BatchDeleteByIdList 批量删除使用id列表
func BatchDeleteByIdList[T model.Model](idList []any) error {
	var t T
	if len(idList) == 0 {
		return fmt.Errorf("idList 不能为空")
	}
	result := DB.Where("id IN ?", idList).Delete(&t)
	return result.Error
}

// Exists 检查记录是否存在
func Exists[T model.Model](conditions map[string]interface{}) (bool, error) {
	var t T
	var count int64
	result := DB.Model(&t).Where(conditions).Count(&count)
	return count > 0, result.Error
}

// Count 计数
func Count[T model.Model]() (int64, error) {
	var t T
	var count int64
	result := DB.Model(&t).Count(&count)
	return count, result.Error
}

// CountByCondition 条件计数
func CountByCondition[T model.Model](conditions map[string]interface{}) (int64, error) {
	var t T
	var count int64
	result := DB.Model(&t).Where(conditions).Count(&count)
	return count, result.Error
}
