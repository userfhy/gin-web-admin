package dictDataService

import (
	"fmt"
	"strings"

	model "gin-web-admin/app/models"
	dictTypeService "gin-web-admin/app/service/v1/dict_type"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/security"
)

var allowedListClass = map[string]struct{}{
	"":        {},
	"default": {},
	"primary": {},
	"success": {},
	"info":    {},
	"warning": {},
	"danger":  {},
}

type DictDataQuery struct {
	Pagination utils.Pagination
	DictType   string
	Label      string
	Status     *int
}

type CreateDictDataStruct struct {
	DictType  string `json:"dictType" binding:"required,max=100"`
	Label     string `json:"label" binding:"required,max=100"`
	Value     string `json:"value" binding:"required,max=100"`
	Status    int    `json:"status" binding:"oneof=0 1"`
	Sort      int    `json:"sort"`
	CSSClass  string `json:"cssClass" binding:"max=100"`
	ListClass string `json:"listClass" binding:"max=50"`
	IsDefault int    `json:"isDefault" binding:"oneof=0 1"`
	Remark    string `json:"remark" binding:"max=500"`
}

type UpdateDictDataStruct struct {
	DictType  string `json:"dictType" binding:"required,max=100"`
	Label     string `json:"label" binding:"required,max=100"`
	Value     string `json:"value" binding:"required,max=100"`
	Status    int    `json:"status" binding:"oneof=0 1"`
	Sort      int    `json:"sort"`
	CSSClass  string `json:"cssClass" binding:"max=100"`
	ListClass string `json:"listClass" binding:"max=50"`
	IsDefault int    `json:"isDefault" binding:"oneof=0 1"`
	Remark    string `json:"remark" binding:"max=500"`
}

type Service struct {
	store *data.Store
}

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetDictDataList(query DictDataQuery) (utils.PageResult, error) {
	if cached, ok, err := getCachedDictDataPage(query); err == nil && ok {
		return cached, nil
	}
	list, total, err := model.GetDictDataList(query.Pagination, query.DictType, query.Label, query.Status)
	if err != nil {
		return utils.PageResult{}, err
	}
	return query.Pagination.Result(list, total), nil
}

func (s *Service) CreateDictData(payload CreateDictDataStruct) error {
	row, err := buildDictDataModel(payload.DictType, payload.Label, payload.Value, payload.Status, payload.Sort, payload.CSSClass, payload.ListClass, payload.IsDefault, payload.Remark)
	if err != nil {
		return err
	}
	if err := model.CreateDictData(row); err != nil {
		return err
	}
	return dictTypeService.NewService(s.store).RefreshDictCache()
}

func (s *Service) UpdateDictData(id int, payload UpdateDictDataStruct) error {
	old, err := model.GetDictDataByID(id)
	if err != nil {
		return err
	}
	if old == nil {
		return fmt.Errorf("dict data not found")
	}

	row, err := buildDictDataModel(payload.DictType, payload.Label, payload.Value, payload.Status, payload.Sort, payload.CSSClass, payload.ListClass, payload.IsDefault, payload.Remark)
	if err != nil {
		return err
	}
	if err := model.UpdateDictData(id, map[string]any{
		"dict_type":  row.DictType,
		"label":      row.Label,
		"value":      row.Value,
		"status":     row.Status,
		"sort":       row.Sort,
		"css_class":  row.CSSClass,
		"list_class": row.ListClass,
		"is_default": row.IsDefault,
		"remark":     row.Remark,
	}); err != nil {
		return err
	}
	return dictTypeService.NewService(s.store).RefreshDictCache()
}

func (s *Service) DeleteDictData(id int) error {
	row, err := model.GetDictDataByID(id)
	if err != nil {
		return err
	}
	if row == nil {
		return fmt.Errorf("dict data not found")
	}
	if err := model.DeleteDictData(id); err != nil {
		return err
	}
	return dictTypeService.NewService(s.store).RefreshDictCache()
}

func buildDictDataModel(dictType, label, value string, status, sort int, cssClass, listClass string, isDefault int, remark string) (model.DictData, error) {
	dictType = strings.TrimSpace(strings.ToLower(dictType))
	if dictType == "" {
		return model.DictData{}, fmt.Errorf("dict type is required")
	}
	dictTypeRow, err := model.GetDictTypeByType(dictType)
	if err != nil {
		return model.DictData{}, err
	}
	if dictTypeRow == nil {
		return model.DictData{}, fmt.Errorf("dict type not found")
	}

	label = security.SanitizePlainText(label, 100)
	value = security.SanitizePlainText(value, 100)
	cssClass = security.SanitizePlainText(cssClass, 100)
	listClass = strings.TrimSpace(strings.ToLower(listClass))
	remark = security.SanitizePlainText(remark, 500)

	if label == "" {
		return model.DictData{}, fmt.Errorf("label is required")
	}
	if value == "" {
		return model.DictData{}, fmt.Errorf("value is required")
	}
	if _, ok := allowedListClass[listClass]; !ok {
		return model.DictData{}, fmt.Errorf("list class is invalid")
	}

	return model.DictData{
		DictType:  dictType,
		Label:     label,
		Value:     value,
		Status:    status,
		Sort:      sort,
		CSSClass:  cssClass,
		ListClass: listClass,
		IsDefault: isDefault,
		Remark:    remark,
	}, nil
}

func getCachedDictDataPage(query DictDataQuery) (utils.PageResult, bool, error) {
	dictType := strings.TrimSpace(strings.ToLower(query.DictType))
	if dictType == "" || strings.TrimSpace(query.Label) != "" {
		return utils.PageResult{}, false, nil
	}
	if query.Status == nil || *query.Status != 1 {
		return utils.PageResult{}, false, nil
	}

	cached, found, err := dictTypeService.LoadCachedDictData(dictType)
	if err != nil || !found {
		return utils.PageResult{}, false, err
	}

	page := query.Pagination.Clone()
	total := int64(len(cached))
	if total == 0 {
		return page.Result([]dictTypeService.CachedDictItem{}, 0), true, nil
	}

	offset := page.Offset()
	if offset >= int(total) {
		return page.Result([]dictTypeService.CachedDictItem{}, total), true, nil
	}

	end := offset + page.Limit()
	if end > int(total) {
		end = int(total)
	}
	return page.Result(cached[offset:end], total), true, nil
}
