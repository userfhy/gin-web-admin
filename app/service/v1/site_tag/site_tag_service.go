package siteTagService

import (
	"fmt"
	"regexp"
	"strings"

	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"gin-web-admin/utils/security"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9\-_/]*$`)

type SiteTagQuery struct {
	Pagination utils.Pagination
	Keyword    string
	Status     *int
}

type SiteTagVO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Status       int    `json:"status"`
	Sort         int    `json:"sort"`
	ContentCount int64  `json:"contentCount"`
	CreatedAt    any    `json:"createdAt"`
	UpdatedAt    any    `json:"updatedAt"`
}

type CreateSiteTagStruct struct {
	Name   string `json:"name" binding:"required,max=80"`
	Slug   string `json:"slug" binding:"required,max=120"`
	Status int    `json:"status" binding:"oneof=0 1"`
	Sort   int    `json:"sort"`
}

type UpdateSiteTagStruct struct {
	Name   string `json:"name" binding:"required,max=80"`
	Slug   string `json:"slug" binding:"required,max=120"`
	Status int    `json:"status" binding:"oneof=0 1"`
	Sort   int    `json:"sort"`
}

func GetSiteTagList(query SiteTagQuery) (utils.PageResult, error) {
	list, total, err := model.GetSiteTagList(query.Pagination, query.Keyword, query.Status)
	if err != nil {
		return utils.PageResult{}, err
	}
	ids := make([]int, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	countMap, err := model.CountSiteContentByTagIDs(ids)
	if err != nil {
		return utils.PageResult{}, err
	}
	vos := make([]SiteTagVO, 0, len(list))
	for _, item := range list {
		vos = append(vos, SiteTagVO{
			ID:           item.ID,
			Name:         item.Name,
			Slug:         item.Slug,
			Status:       item.Status,
			Sort:         item.Sort,
			ContentCount: countMap[item.ID],
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}
	return query.Pagination.Result(vos, total), nil
}

func GetAllSiteTags(status *int) ([]*model.SiteTag, error) {
	return model.GetAllSiteTags(status)
}

func CreateSiteTag(payload CreateSiteTagStruct) error {
	name := security.SanitizePlainText(payload.Name, 80)
	slug := normalizeSlug(payload.Slug)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteTagBySlug(slug, 0)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}

	return model.CreateSiteTag(model.SiteTag{
		Name:   name,
		Slug:   slug,
		Status: payload.Status,
		Sort:   payload.Sort,
	})
}

func UpdateSiteTag(id int, payload UpdateSiteTagStruct) error {
	name := security.SanitizePlainText(payload.Name, 80)
	slug := normalizeSlug(payload.Slug)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateSlug(slug); err != nil {
		return err
	}

	exists, err := model.ExistsSiteTagBySlug(slug, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("slug already exists")
	}

	data := map[string]any{
		"name":   name,
		"slug":   slug,
		"status": payload.Status,
		"sort":   payload.Sort,
	}
	return model.UpdateSiteTag(id, data)
}

func DeleteSiteTag(id int) error {
	inUse, err := model.IsSiteTagInUse(id)
	if err != nil {
		return err
	}
	if inUse {
		return fmt.Errorf("tag in use")
	}
	return model.DeleteSiteTag(id)
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
