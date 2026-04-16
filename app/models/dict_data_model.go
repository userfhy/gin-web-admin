package model

import (
	"errors"
	"strings"
	"time"

	"gin-web-admin/utils"
	"gorm.io/gorm"
)

type DictData struct {
	ID        int       `gorm:"primaryKey;comment:主键ID" json:"id"`
	DictType  string    `gorm:"type:varchar(100);not null;index:idx_dict_data_type;comment:字典类型" json:"dictType"`
	Label     string    `gorm:"type:varchar(100);not null;comment:字典标签" json:"label"`
	Value     string    `gorm:"type:varchar(100);not null;comment:字典键值" json:"value"`
	Status    int       `gorm:"type:int(1);default:1;not null;comment:状态(1启用0停用)" json:"status"`
	Sort      int       `gorm:"default:0;not null;comment:排序值(越小越靠前)" json:"sort"`
	CSSClass  string    `gorm:"column:css_class;type:varchar(100);comment:CSS类名" json:"cssClass"`
	ListClass string    `gorm:"column:list_class;type:varchar(50);comment:回显样式" json:"listClass"`
	IsDefault int       `gorm:"column:is_default;type:int(1);default:0;not null;comment:是否默认(1是0否)" json:"isDefault"`
	Remark    string    `gorm:"type:varchar(500);comment:备注" json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (DictData) TableName() string {
	return TablePrefix + "dict_data"
}

func GetDictDataByID(id int) (*DictData, error) {
	var row DictData
	err := db.First(&row, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

func GetDictDataList(pg utils.Pagination, dictType, label string, status *int) ([]*DictData, int64, error) {
	var (
		list  []*DictData
		total int64
	)

	query := db.Model(&DictData{})
	if dictType = strings.TrimSpace(dictType); dictType != "" {
		query = query.Where("dict_type = ?", dictType)
	}
	if label = strings.TrimSpace(label); label != "" {
		query = query.Where("label LIKE ?", "%"+label+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("sort ASC").Order("id ASC").
		Scopes(pg.Scope()).
		Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return list, total, nil
}

func CreateDictData(row DictData) error {
	return db.Create(&row).Error
}

func UpdateDictData(id int, data map[string]any) error {
	return db.Model(&DictData{}).Where("id = ?", id).Updates(data).Error
}

func DeleteDictData(id int) error {
	return db.Delete(&DictData{}, id).Error
}
