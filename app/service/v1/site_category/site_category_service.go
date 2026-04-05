package siteCategoryService

import (
	"fmt"
	model "gin-web-admin/app/models"
	"gin-web-admin/utils/security"
	"regexp"
	"strings"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9\-_/]*$`)

type SiteCategoryQuery struct {
	PageNum  int
	PageSize int
	Keyword  string
	Status   *int
}

type SiteCategoryVO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	Status       int    `json:"status"`
	Sort         int    `json:"sort"`
	ContentCount int64  `json:"contentCount"`
	CreatedAt    any    `json:"createdAt"`
	UpdatedAt    any    `json:"updatedAt"`
}

type CreateSiteCategoryStruct struct {
	Name        string `json:"name" binding:"required,max=100"`
	Slug        string `json:"slug" binding:"required,max=120"`
	Description string `json:"description" binding:"max=500"`
	Status      int    `json:"status" binding:"oneof=0 1"`
	Sort        int    `json:"sort"`
}

type UpdateSiteCategoryStruct struct {
	Name        string `json:"name" binding:"required,max=100"`
	Slug        string `json:"slug" binding:"required,max=120"`
	Description string `json:"description" binding:"max=500"`
	Status      int    `json:"status" binding:"oneof=0 1"`
	Sort        int    `json:"sort"`
}

func GetSiteCategoryList(query SiteCategoryQuery) (map[string]any, error) {
	list, total, err := model.GetSiteCategoryList(query.PageNum, query.PageSize, query.Keyword, query.Status)
	if err != nil {
		return nil, err
	}
	ids := make([]int, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	countMap, err := model.CountSiteContentByCategoryIDs(ids)
	if err != nil {
		return nil, err
	}
	vos := make([]SiteCategoryVO, 0, len(list))
	for _, item := range list {
		vos = append(vos, SiteCategoryVO{
			ID:           item.ID,
			Name:         item.Name,
			Slug:         item.Slug,
			Description:  item.Description,
			Status:       item.Status,
			Sort:         item.Sort,
			ContentCount: countMap[item.ID],
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}
	return map[string]any{
		"list":        vos,
		"total":       total,
		"currentPage": query.PageNum,
		"pageSize":    query.PageSize,
	}, nil
}

func GetAllSiteCategories(status *int) ([]*model.SiteCategory, error) {
	return model.GetAllSiteCategories(status)
}

func CreateSiteCategory(payload CreateSiteCategoryStruct) error {
	name := security.SanitizePlainText(payload.Name, 100)
	slug := normalizeSlug(payload.Slug)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteCategoryBySlug(slug, 0)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}

	return model.CreateSiteCategory(model.SiteCategory{
		Name:        name,
		Slug:        slug,
		Description: security.SanitizePlainText(payload.Description, 500),
		Status:      payload.Status,
		Sort:        payload.Sort,
	})
}

func UpdateSiteCategory(id int, payload UpdateSiteCategoryStruct) error {
	name := security.SanitizePlainText(payload.Name, 100)
	slug := normalizeSlug(payload.Slug)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteCategoryBySlug(slug, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}

	data := map[string]any{
		"name":        name,
		"slug":        slug,
		"description": security.SanitizePlainText(payload.Description, 500),
		"status":      payload.Status,
		"sort":        payload.Sort,
	}
	return model.UpdateSiteCategory(id, data)
}

func DeleteSiteCategory(id int) error {
	inUse, err := model.IsSiteCategoryInUse(id)
	if err != nil {
		return err
	}
	if inUse {
		return fmt.Errorf("category in use")
	}
	return model.DeleteSiteCategory(id)
}

func validateSlug(slug string) error {
	if slug == "" {
		return fmt.Errorf("slug is required")
	}
	if len(slug) > 120 {
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
