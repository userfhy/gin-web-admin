package service

import (
	model "gin-web-admin/app/models"
	sysService "gin-web-admin/app/service/v1/sys"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
)

type Service struct {
	store *data.Store
}

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateMenu(menu model.Menu) error {
	if err := model.CreateMenu(menu); err != nil {
		return err
	}
	sysService.InvalidateRouteCache()
	return nil
}

func (s *Service) UpdateMenu(id int, data map[string]any) error {
	if err := model.UpdateMenu(id, data); err != nil {
		return err
	}
	sysService.InvalidateRouteCache()
	return nil
}

func (s *Service) DeleteMenu(id int) error {
	if err := model.DeleteMenu(id); err != nil {
		return err
	}
	sysService.InvalidateRouteCache()
	return nil
}

func (s *Service) GetMenu(where map[string]any) (*model.Menu, error) {
	return model.GetMenu(where)
}

func (s *Service) GetMenuList(pg utils.Pagination, where map[string]any) ([]*model.Menu, error) {
	return model.GetMenuList(pg, where)
}

func (s *Service) GetAllMenus(where map[string]any) ([]*model.Menu, error) {
	return model.GetAllMenus(where)
}
