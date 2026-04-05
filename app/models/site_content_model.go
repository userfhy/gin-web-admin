package model

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SiteContent struct {
	ID             int        `gorm:"primaryKey;comment:主键ID" json:"id"`
	Title          string     `gorm:"type:varchar(150);not null;comment:内容标题" json:"title"`
	Slug           string     `gorm:"type:varchar(150);not null;uniqueIndex:idx_site_content_slug;comment:页面标识" json:"slug"`
	Summary        string     `gorm:"type:varchar(500);comment:摘要" json:"summary"`
	Cover          string     `gorm:"type:varchar(500);comment:封面图" json:"cover"`
	Content        string     `gorm:"type:longtext;comment:正文内容" json:"content"`
	SeoKeywords    string     `gorm:"type:varchar(255);comment:SEO关键词" json:"seoKeywords"`
	SeoDescription string     `gorm:"type:varchar(500);comment:SEO描述" json:"seoDescription"`
	Status         int        `gorm:"type:int(1);default:1;not null;comment:状态(1发布0草稿)" json:"status"`
	PublishedAt    *time.Time `gorm:"comment:发布时间" json:"publishedAt"`
	Sort           int        `gorm:"default:0;not null;comment:排序值(越小越靠前)" json:"sort"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

func (SiteContent) TableName() string {
	return TablePrefix + "site_content"
}

func GetSiteContentByID(id int) (*SiteContent, error) {
	var content SiteContent
	err := db.First(&content, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &content, nil
}

func GetSiteContentBySlug(slug string) (*SiteContent, error) {
	var content SiteContent
	err := db.Where("slug = ?", slug).First(&content).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &content, nil
}

func GetSiteContentList(pageNum, pageSize int, keyword string, status *int, categoryID *int, tagID *int) ([]*SiteContent, int64, error) {
	var (
		list  []*SiteContent
		total int64
	)

	query := db.Model(&SiteContent{})
	if categoryID != nil {
		query = query.Joins("JOIN "+TablePrefix+"site_content_category scc ON scc.content_id = "+TablePrefix+"site_content.id").
			Where("scc.category_id = ?", *categoryID)
	}
	if tagID != nil {
		query = query.Joins("JOIN "+TablePrefix+"site_content_tag sct ON sct.content_id = "+TablePrefix+"site_content.id").
			Where("sct.tag_id = ?", *tagID)
	}

	if keyword = strings.TrimSpace(keyword); keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("title LIKE ? OR slug LIKE ?", like, like)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("sort ASC").
		Order("id DESC").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return list, total, nil
}

func GetPublishedSiteContentList(pageNum, pageSize int, keyword string, categoryID *int, tagID *int) ([]*SiteContent, int64, error) {
	status := 1
	return GetSiteContentList(pageNum, pageSize, keyword, &status, categoryID, tagID)
}

func ExistsSiteContentBySlug(slug string, excludeID int) (bool, error) {
	var count int64
	query := db.Model(&SiteContent{}).Where("slug = ?", slug)
	if excludeID > 0 {
		query = query.Where("id <> ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateSiteContent(content SiteContent) error {
	return db.Create(&content).Error
}

func CreateSiteContentTx(tx *gorm.DB, content *SiteContent) error {
	if tx == nil {
		tx = db
	}
	return tx.Create(content).Error
}

func UpdateSiteContent(id int, data map[string]any) error {
	return db.Model(&SiteContent{}).Where("id = ?", id).Updates(data).Error
}

func UpdateSiteContentStatus(id int, status int, publishedAt *time.Time) error {
	data := map[string]any{
		"status": status,
	}
	if publishedAt != nil {
		data["published_at"] = *publishedAt
	}
	return db.Model(&SiteContent{}).Where("id = ?", id).Updates(data).Error
}

func DeleteSiteContent(id int) error {
	return db.Delete(&SiteContent{}, id).Error
}
