package model

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SiteTag struct {
	ID        int       `gorm:"primaryKey;comment:主键ID" json:"id"`
	Name      string    `gorm:"type:varchar(80);not null;comment:标签名" json:"name"`
	Slug      string    `gorm:"type:varchar(120);not null;uniqueIndex:idx_site_tag_slug;comment:标签标识" json:"slug"`
	Status    int       `gorm:"type:int(1);default:1;not null;comment:状态(1启用0停用)" json:"status"`
	Sort      int       `gorm:"default:0;not null;comment:排序值(越小越靠前)" json:"sort"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (SiteTag) TableName() string {
	return TablePrefix + "site_tag"
}

func GetSiteTagByID(id int) (*SiteTag, error) {
	var tag SiteTag
	err := db.First(&tag, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func GetSiteTagBySlug(slug string) (*SiteTag, error) {
	var tag SiteTag
	err := db.Where("slug = ?", slug).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func GetSiteTagList(pageNum, pageSize int, keyword string, status *int) ([]*SiteTag, int64, error) {
	var (
		list  []*SiteTag
		total int64
	)

	query := db.Model(&SiteTag{})
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

func GetAllSiteTags(status *int) ([]*SiteTag, error) {
	var list []*SiteTag
	query := db.Model(&SiteTag{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	err := query.Order("sort ASC").Order("id ASC").Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func ExistsSiteTagBySlug(slug string, excludeID int) (bool, error) {
	var count int64
	query := db.Model(&SiteTag{}).Where("slug = ?", slug)
	if excludeID > 0 {
		query = query.Where("id <> ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateSiteTag(tag SiteTag) error {
	return db.Create(&tag).Error
}

func UpdateSiteTag(id int, data map[string]any) error {
	return db.Model(&SiteTag{}).Where("id = ?", id).Updates(data).Error
}

func DeleteSiteTag(id int) error {
	return db.Delete(&SiteTag{}, id).Error
}

func CountSiteTagsByIDs(ids []int) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	var count int64
	err := db.Model(&SiteTag{}).Where("id IN ?", ids).Count(&count).Error
	return count, err
}

func GetSiteTagsByIDs(ids []int) ([]*SiteTag, error) {
	if len(ids) == 0 {
		return []*SiteTag{}, nil
	}
	var list []*SiteTag
	err := db.Where("id IN ?", ids).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func CountSiteContentByTagIDs(ids []int) (map[int]int64, error) {
	result := make(map[int]int64, len(ids))
	if len(ids) == 0 {
		return result, nil
	}

	type row struct {
		TagID int   `gorm:"column:tag_id"`
		Total int64 `gorm:"column:total"`
	}
	var rows []row
	err := db.Model(&SiteContentTag{}).
		Select("tag_id, COUNT(*) AS total").
		Where("tag_id IN ?", ids).
		Group("tag_id").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	for _, r := range rows {
		result[r.TagID] = r.Total
	}
	return result, nil
}
