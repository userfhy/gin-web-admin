package siteContentService

import (
	"fmt"
	model "gin-web-admin/app/models"
	"regexp"
	"sort"
	"strings"

	"gorm.io/gorm"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9\-_/]*$`)

type SiteContentQuery struct {
	PageNum    int
	PageSize   int
	Keyword    string
	Status     *int
	CategoryID *int
}

type SiteContentVO struct {
	ID             int      `json:"id"`
	Title          string   `json:"title"`
	Slug           string   `json:"slug"`
	Summary        string   `json:"summary"`
	Cover          string   `json:"cover"`
	Content        string   `json:"content"`
	SeoKeywords    string   `json:"seoKeywords"`
	SeoDescription string   `json:"seoDescription"`
	Status         int      `json:"status"`
	Sort           int      `json:"sort"`
	CreatedAt      any      `json:"createdAt"`
	UpdatedAt      any      `json:"updatedAt"`
	CategoryIDs    []int    `json:"categoryIds"`
	CategoryNames  []string `json:"categoryNames"`
}

type CreateSiteContentStruct struct {
	Title          string `json:"title" binding:"required,max=150"`
	Slug           string `json:"slug" binding:"required,max=150"`
	Summary        string `json:"summary" binding:"max=500"`
	Cover          string `json:"cover" binding:"max=500"`
	Content        string `json:"content"`
	SeoKeywords    string `json:"seoKeywords" binding:"max=255"`
	SeoDescription string `json:"seoDescription" binding:"max=500"`
	Status         int    `json:"status" binding:"oneof=0 1"`
	Sort           int    `json:"sort"`
	CategoryIDs    []int  `json:"categoryIds"`
}

type UpdateSiteContentStruct struct {
	Title          string `json:"title" binding:"required,max=150"`
	Slug           string `json:"slug" binding:"required,max=150"`
	Summary        string `json:"summary" binding:"max=500"`
	Cover          string `json:"cover" binding:"max=500"`
	Content        string `json:"content"`
	SeoKeywords    string `json:"seoKeywords" binding:"max=255"`
	SeoDescription string `json:"seoDescription" binding:"max=500"`
	Status         int    `json:"status" binding:"oneof=0 1"`
	Sort           int    `json:"sort"`
	CategoryIDs    []int  `json:"categoryIds"`
}

func GetSiteContentList(query SiteContentQuery) (map[string]any, error) {
	list, total, err := model.GetSiteContentList(query.PageNum, query.PageSize, query.Keyword, query.Status, query.CategoryID)
	if err != nil {
		return nil, err
	}
	contentVOs, err := toSiteContentVOList(list)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"list":        contentVOs,
		"total":       total,
		"currentPage": query.PageNum,
		"pageSize":    query.PageSize,
	}, nil
}

func GetSiteContentDetail(id int) (*SiteContentVO, error) {
	row, err := model.GetSiteContentByID(id)
	if err != nil || row == nil {
		return nil, err
	}
	list, err := toSiteContentVOList([]*model.SiteContent{row})
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func CreateSiteContent(payload CreateSiteContentStruct) error {
	title := strings.TrimSpace(payload.Title)
	slug := normalizeSlug(payload.Slug)
	categoryIDs := uniqueSortedIDs(payload.CategoryIDs)

	if title == "" {
		return fmt.Errorf("title is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteContentBySlug(slug, 0)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}
	if err := validateCategoryIDs(categoryIDs); err != nil {
		return err
	}

	content := &model.SiteContent{
		Title:          title,
		Slug:           slug,
		Summary:        strings.TrimSpace(payload.Summary),
		Cover:          strings.TrimSpace(payload.Cover),
		Content:        payload.Content,
		SeoKeywords:    strings.TrimSpace(payload.SeoKeywords),
		SeoDescription: strings.TrimSpace(payload.SeoDescription),
		Status:         payload.Status,
		Sort:           payload.Sort,
	}

	return model.DB().Transaction(func(tx *gorm.DB) error {
		if err := model.CreateSiteContentTx(tx, content); err != nil {
			return err
		}
		return model.ReplaceSiteContentCategoriesTx(tx, content.ID, categoryIDs)
	})
}

func UpdateSiteContent(id int, payload UpdateSiteContentStruct) error {
	title := strings.TrimSpace(payload.Title)
	slug := normalizeSlug(payload.Slug)
	categoryIDs := uniqueSortedIDs(payload.CategoryIDs)

	if title == "" {
		return fmt.Errorf("title is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteContentBySlug(slug, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}
	if err := validateCategoryIDs(categoryIDs); err != nil {
		return err
	}

	data := map[string]any{
		"title":           title,
		"slug":            slug,
		"summary":         strings.TrimSpace(payload.Summary),
		"cover":           strings.TrimSpace(payload.Cover),
		"content":         payload.Content,
		"seo_keywords":    strings.TrimSpace(payload.SeoKeywords),
		"seo_description": strings.TrimSpace(payload.SeoDescription),
		"status":          payload.Status,
		"sort":            payload.Sort,
	}

	return model.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.SiteContent{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return err
		}
		return model.ReplaceSiteContentCategoriesTx(tx, id, categoryIDs)
	})
}

func DeleteSiteContent(id int) error {
	return model.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("content_id = ?", id).Delete(&model.SiteContentCategory{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.SiteContent{}, id).Error
	})
}

func toSiteContentVOList(list []*model.SiteContent) ([]SiteContentVO, error) {
	result := make([]SiteContentVO, 0, len(list))
	if len(list) == 0 {
		return result, nil
	}

	ids := make([]int, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}

	categoryIDsMap, err := model.GetCategoryIDsByContentIDs(ids)
	if err != nil {
		return nil, err
	}

	allCategoryIDSet := make(map[int]struct{})
	for _, cids := range categoryIDsMap {
		for _, cid := range cids {
			allCategoryIDSet[cid] = struct{}{}
		}
	}
	allCategoryIDs := make([]int, 0, len(allCategoryIDSet))
	for cid := range allCategoryIDSet {
		allCategoryIDs = append(allCategoryIDs, cid)
	}

	categoryNameMap := make(map[int]string, len(allCategoryIDs))
	if len(allCategoryIDs) > 0 {
		categories, err := model.GetSiteCategoriesByIDs(allCategoryIDs)
		if err != nil {
			return nil, err
		}
		for _, category := range categories {
			categoryNameMap[category.ID] = category.Name
		}
	}

	for _, item := range list {
		cids := uniqueSortedIDs(categoryIDsMap[item.ID])
		cnames := make([]string, 0, len(cids))
		for _, cid := range cids {
			if name, ok := categoryNameMap[cid]; ok {
				cnames = append(cnames, name)
			}
		}
		result = append(result, SiteContentVO{
			ID:             item.ID,
			Title:          item.Title,
			Slug:           item.Slug,
			Summary:        item.Summary,
			Cover:          item.Cover,
			Content:        item.Content,
			SeoKeywords:    item.SeoKeywords,
			SeoDescription: item.SeoDescription,
			Status:         item.Status,
			Sort:           item.Sort,
			CreatedAt:      item.CreatedAt,
			UpdatedAt:      item.UpdatedAt,
			CategoryIDs:    cids,
			CategoryNames:  cnames,
		})
	}
	return result, nil
}

func validateCategoryIDs(categoryIDs []int) error {
	if len(categoryIDs) == 0 {
		return nil
	}
	count, err := model.CountSiteCategoriesByIDs(categoryIDs)
	if err != nil {
		return err
	}
	if int(count) != len(categoryIDs) {
		return fmt.Errorf("categoryIds contains invalid id")
	}
	return nil
}

func uniqueSortedIDs(ids []int) []int {
	if len(ids) == 0 {
		return []int{}
	}
	idSet := make(map[int]struct{}, len(ids))
	for _, id := range ids {
		if id > 0 {
			idSet[id] = struct{}{}
		}
	}
	result := make([]int, 0, len(idSet))
	for id := range idSet {
		result = append(result, id)
	}
	sort.Ints(result)
	return result
}

func validateSlug(slug string) error {
	if slug == "" {
		return fmt.Errorf("slug is required")
	}
	if len(slug) > 150 {
		return fmt.Errorf("slug too long")
	}
	if !slugPattern.MatchString(slug) {
		return fmt.Errorf("slug format is invalid")
	}
	return nil
}

func normalizeSlug(slug string) string {
	return strings.Trim(strings.ToLower(strings.TrimSpace(slug)), "/")
}
