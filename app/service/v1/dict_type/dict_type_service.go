package dictTypeService

import (
	"fmt"
	"regexp"
	"strings"

	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/security"
	"gorm.io/gorm"
)

var dictTypePattern = regexp.MustCompile(`^[a-z][a-z0-9_:-]*$`)

type DictTypeQuery struct {
	Pagination utils.Pagination
	Name       string
	Type       string
	Status     *int
}

type DictTypeVO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Status    int    `json:"status"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	DataCount int64  `json:"dataCount"`
	CreatedAt any    `json:"createdAt"`
	UpdatedAt any    `json:"updatedAt"`
}

type CreateDictTypeStruct struct {
	Name   string `json:"name" binding:"required,max=100"`
	Type   string `json:"type" binding:"required,max=100"`
	Status int    `json:"status" binding:"oneof=0 1"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark" binding:"max=500"`
}

type UpdateDictTypeStruct struct {
	Name   string `json:"name" binding:"required,max=100"`
	Type   string `json:"type" binding:"required,max=100"`
	Status int    `json:"status" binding:"oneof=0 1"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark" binding:"max=500"`
}

type Service struct {
	store *data.Store
}

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetDictTypeList(query DictTypeQuery) (utils.PageResult, error) {
	list, total, err := model.GetDictTypeList(query.Pagination, query.Name, query.Type, query.Status)
	if err != nil {
		return utils.PageResult{}, err
	}

	types := make([]string, 0, len(list))
	for _, item := range list {
		types = append(types, item.Type)
	}
	countMap, err := model.CountDictDataByTypes(types)
	if err != nil {
		return utils.PageResult{}, err
	}

	vos := make([]DictTypeVO, 0, len(list))
	for _, item := range list {
		vos = append(vos, DictTypeVO{
			ID:        item.ID,
			Name:      item.Name,
			Type:      item.Type,
			Status:    item.Status,
			Sort:      item.Sort,
			Remark:    item.Remark,
			DataCount: countMap[item.Type],
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return query.Pagination.Result(vos, total), nil
}

func (s *Service) GetAllDictTypes(status *int) ([]*model.DictType, error) {
	return model.GetAllDictTypes(status)
}

func (s *Service) RefreshDictCache() error {
	return refreshDictCache()
}

func (s *Service) CreateDictType(payload CreateDictTypeStruct) error {
	name := security.SanitizePlainText(payload.Name, 100)
	dictType := normalizeDictType(payload.Type)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateDictType(dictType); err != nil {
		return err
	}

	exists, err := model.ExistsDictTypeByType(dictType, 0)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("dict type already exists")
	}

	if err := model.CreateDictType(model.DictType{
		Name:   name,
		Type:   dictType,
		Status: payload.Status,
		Sort:   payload.Sort,
		Remark: security.SanitizePlainText(payload.Remark, 500),
	}); err != nil {
		return err
	}
	return refreshDictCache(dictType)
}

func (s *Service) UpdateDictType(id int, payload UpdateDictTypeStruct) error {
	name := security.SanitizePlainText(payload.Name, 100)
	dictType := normalizeDictType(payload.Type)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if err := validateDictType(dictType); err != nil {
		return err
	}

	old, err := model.GetDictTypeByID(id)
	if err != nil {
		return err
	}
	if old == nil {
		return fmt.Errorf("dict type not found")
	}

	exists, err := model.ExistsDictTypeByType(dictType, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("dict type already exists")
	}

	if err := model.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.DictType{}).Where("id = ?", id).Updates(map[string]any{
			"name":   name,
			"type":   dictType,
			"status": payload.Status,
			"sort":   payload.Sort,
			"remark": security.SanitizePlainText(payload.Remark, 500),
		}).Error; err != nil {
			return err
		}

		if old.Type != dictType {
			if err := tx.Model(&model.DictData{}).Where("dict_type = ?", old.Type).Update("dict_type", dictType).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	if old.Type != dictType {
		return refreshDictCache(old.Type, dictType)
	}
	return refreshDictCache(dictType)
}

func (s *Service) DeleteDictType(id int) error {
	row, err := model.GetDictTypeByID(id)
	if err != nil {
		return err
	}
	if row == nil {
		return fmt.Errorf("dict type not found")
	}

	inUse, err := model.IsDictTypeInUse(row.Type)
	if err != nil {
		return err
	}
	if inUse {
		return fmt.Errorf("dict type in use")
	}
	if err := model.DeleteDictType(id); err != nil {
		return err
	}
	return refreshDictCache(row.Type)
}

func validateDictType(dictType string) error {
	if dictType == "" {
		return fmt.Errorf("dict type is required")
	}
	if len(dictType) > 100 {
		return fmt.Errorf("dict type too long")
	}
	if !dictTypePattern.MatchString(dictType) {
		return fmt.Errorf("dict type format is invalid")
	}
	return nil
}

func normalizeDictType(dictType string) string {
	return strings.TrimSpace(strings.ToLower(dictType))
}
