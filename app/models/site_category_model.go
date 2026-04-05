package model

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SiteCategory struct {
	ID          int       `gorm:"primaryKey;comment:主键ID" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null;comment:分类名称" json:"name"`
	Slug        string    `gorm:"type:varchar(120);not null;uniqueIndex:idx_site_category_slug;comment:分类标识" json:"slug"`
	Description string    `gorm:"type:varchar(500);comment:分类描述" json:"description"`
	Status      int       `gorm:"type:int(1);default:1;not null;comment:状态(1启用0停用)" json:"status"`
	Sort        int       `gorm:"default:0;not null;comment:排序值(越小越靠前)" json:"sort"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (SiteCategory) TableName() string {
	return TablePrefix + "site_category"
}

func GetSiteCategoryByID(id int) (*SiteCategory, error) {
	var category SiteCategory
	err := db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func GetSiteCategoryList(pageNum, pageSize int, keyword string, status *int) ([]*SiteCategory, int64, error) {
	var (
		list  []*SiteCategory
		total int64
	)

	query := db.Model(&SiteCategory{})
	if keyword = strings.TrimSpace(keyword); keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR slug LIKE ?", like, like)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("sort ASC").Order("id DESC").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return list, total, nil
}

func GetAllSiteCategories(status *int) ([]*SiteCategory, error) {
	var list []*SiteCategory
	query := db.Model(&SiteCategory{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	err := query.Order("sort ASC").Order("id ASC").Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func ExistsSiteCategoryBySlug(slug string, excludeID int) (bool, error) {
	var count int64
	query := db.Model(&SiteCategory{}).Where("slug = ?", slug)
	if excludeID > 0 {
		query = query.Where("id <> ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateSiteCategory(category SiteCategory) error {
	return db.Create(&category).Error
}

func UpdateSiteCategory(id int, data map[string]any) error {
	return db.Model(&SiteCategory{}).Where("id = ?", id).Updates(data).Error
}

func DeleteSiteCategory(id int) error {
	return db.Delete(&SiteCategory{}, id).Error
}

func CountSiteCategoriesByIDs(ids []int) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	var count int64
	err := db.Model(&SiteCategory{}).Where("id IN ?", ids).Count(&count).Error
	return count, err
}

func GetSiteCategoriesByIDs(ids []int) ([]*SiteCategory, error) {
	if len(ids) == 0 {
		return []*SiteCategory{}, nil
	}
	var list []*SiteCategory
	err := db.Where("id IN ?", ids).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}
