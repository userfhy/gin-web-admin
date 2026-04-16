package model

import (
	"errors"
	"strings"
	"time"

	"gin-web-admin/utils"
	"gorm.io/gorm"
)

type DictType struct {
	ID        int       `gorm:"primaryKey;comment:主键ID" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;comment:字典名称" json:"name"`
	Type      string    `gorm:"column:type;type:varchar(100);not null;uniqueIndex:idx_dict_type;comment:字典类型" json:"type"`
	Status    int       `gorm:"type:int(1);default:1;not null;comment:状态(1启用0停用)" json:"status"`
	Sort      int       `gorm:"default:0;not null;comment:排序值(越小越靠前)" json:"sort"`
	Remark    string    `gorm:"type:varchar(500);comment:备注" json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (DictType) TableName() string {
	return TablePrefix + "dict_type"
}

func GetDictTypeByID(id int) (*DictType, error) {
	var row DictType
	err := db.First(&row, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

func GetDictTypeByType(dictType string) (*DictType, error) {
	var row DictType
	err := db.Where("type = ?", dictType).First(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

func GetDictTypeList(pg utils.Pagination, name, dictType string, status *int) ([]*DictType, int64, error) {
	var (
		list  []*DictType
		total int64
	)

	query := db.Model(&DictType{})
	if name = strings.TrimSpace(name); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if dictType = strings.TrimSpace(dictType); dictType != "" {
		query = query.Where("type LIKE ?", "%"+dictType+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("sort ASC").Order("id DESC").
		Scopes(pg.Scope()).
		Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return list, total, nil
}

func GetAllDictTypes(status *int) ([]*DictType, error) {
	var list []*DictType
	query := db.Model(&DictType{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	err := query.Order("sort ASC").Order("id ASC").Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func ExistsDictTypeByType(dictType string, excludeID int) (bool, error) {
	var count int64
	query := db.Model(&DictType{}).Where("type = ?", dictType)
	if excludeID > 0 {
		query = query.Where("id <> ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDictType(row DictType) error {
	return db.Create(&row).Error
}

func UpdateDictType(id int, data map[string]any) error {
	return db.Model(&DictType{}).Where("id = ?", id).Updates(data).Error
}

func DeleteDictType(id int) error {
	return db.Delete(&DictType{}, id).Error
}

func CountDictDataByTypes(types []string) (map[string]int64, error) {
	result := make(map[string]int64, len(types))
	if len(types) == 0 {
		return result, nil
	}

	type row struct {
		DictType string `gorm:"column:dict_type"`
		Total    int64  `gorm:"column:total"`
	}
	var rows []row
	err := db.Model(&DictData{}).
		Select("dict_type, COUNT(*) AS total").
		Where("dict_type IN ?", types).
		Group("dict_type").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	for _, item := range rows {
		result[item.DictType] = item.Total
	}
	return result, nil
}

func IsDictTypeInUse(dictType string) (bool, error) {
	var count int64
	err := db.Model(&DictData{}).Where("dict_type = ?", dictType).Count(&count).Error
	return count > 0, err
}
