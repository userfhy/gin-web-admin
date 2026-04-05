package model

import "gorm.io/gorm"

type SiteContentCategory struct {
	ContentID  int `gorm:"column:content_id;primaryKey" json:"contentId"`
	CategoryID int `gorm:"column:category_id;primaryKey" json:"categoryId"`
}

func (SiteContentCategory) TableName() string {
	return TablePrefix + "site_content_category"
}

func ReplaceSiteContentCategoriesTx(tx *gorm.DB, contentID int, categoryIDs []int) error {
	if tx == nil {
		tx = db
	}
	if err := tx.Where("content_id = ?", contentID).Delete(&SiteContentCategory{}).Error; err != nil {
		return err
	}
	if len(categoryIDs) == 0 {
		return nil
	}
	rows := make([]SiteContentCategory, 0, len(categoryIDs))
	for _, cid := range categoryIDs {
		rows = append(rows, SiteContentCategory{ContentID: contentID, CategoryID: cid})
	}
	return tx.Create(&rows).Error
}

func GetCategoryIDsByContentIDs(contentIDs []int) (map[int][]int, error) {
	result := make(map[int][]int, len(contentIDs))
	if len(contentIDs) == 0 {
		return result, nil
	}

	var rows []SiteContentCategory
	if err := db.Where("content_id IN ?", contentIDs).Find(&rows).Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		result[row.ContentID] = append(result[row.ContentID], row.CategoryID)
	}
	return result, nil
}

func IsSiteCategoryInUse(categoryID int) (bool, error) {
	var count int64
	err := db.Model(&SiteContentCategory{}).Where("category_id = ?", categoryID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
