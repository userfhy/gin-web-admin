package model

import "gorm.io/gorm"

type SiteContentTag struct {
	ContentID int `gorm:"column:content_id;primaryKey" json:"contentId"`
	TagID     int `gorm:"column:tag_id;primaryKey" json:"tagId"`
}

func (SiteContentTag) TableName() string {
	return TablePrefix + "site_content_tag"
}

func ReplaceSiteContentTagsTx(tx *gorm.DB, contentID int, tagIDs []int) error {
	if tx == nil {
		tx = db
	}
	if err := tx.Where("content_id = ?", contentID).Delete(&SiteContentTag{}).Error; err != nil {
		return err
	}
	if len(tagIDs) == 0 {
		return nil
	}
	rows := make([]SiteContentTag, 0, len(tagIDs))
	for _, tid := range tagIDs {
		rows = append(rows, SiteContentTag{ContentID: contentID, TagID: tid})
	}
	return tx.Create(&rows).Error
}

func GetTagIDsByContentIDs(contentIDs []int) (map[int][]int, error) {
	result := make(map[int][]int, len(contentIDs))
	if len(contentIDs) == 0 {
		return result, nil
	}
	var rows []SiteContentTag
	if err := db.Where("content_id IN ?", contentIDs).Find(&rows).Error; err != nil {
		return nil, err
	}
	for _, row := range rows {
		result[row.ContentID] = append(result[row.ContentID], row.TagID)
	}
	return result, nil
}

func IsSiteTagInUse(tagID int) (bool, error) {
	var count int64
	err := db.Model(&SiteContentTag{}).Where("tag_id = ?", tagID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
