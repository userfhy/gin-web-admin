package model

import (
	"errors"

	"gorm.io/gorm"
)

type Menu struct {
	ID              int    `gorm:"primaryKey;comment:主键ID" json:"id"`                      // 主键ID
	ParentID        int    `gorm:"comment:父菜单ID" json:"parentId"`                          // 父菜单ID
	MenuType        int    `gorm:"comment:菜单类型（0:目录 1:菜单 2:按钮）" json:"menuType"`           // 菜单类型
	Title           string `gorm:"type:varchar(100);comment:菜单标题" json:"title"`            // 菜单标题
	Name            string `gorm:"type:varchar(100);comment:菜单名称（唯一值）" json:"name"`        // 菜单名称
	Path            string `gorm:"type:varchar(255);comment:路由地址" json:"path"`             // 路由地址
	Component       string `gorm:"type:varchar(255);comment:组件路径" json:"component"`        // 组件路径
	Rank            *int   `gorm:"comment:排序序号" json:"rank"`                               // 排序序号
	Redirect        string `gorm:"type:varchar(255);comment:重定向地址" json:"redirect"`        // 重定向地址
	Icon            string `gorm:"type:varchar(100);comment:图标" json:"icon"`               // 图标
	ExtraIcon       string `gorm:"type:varchar(100);comment:额外图标" json:"extraIcon"`        // 额外图标
	EnterTransition string `gorm:"type:varchar(100);comment:进入动画" json:"enterTransition"`  // 进入动画
	LeaveTransition string `gorm:"type:varchar(100);comment:离开动画" json:"leaveTransition"`  // 离开动画
	ActivePath      string `gorm:"type:varchar(255);comment:激活路径" json:"activePath"`       // 激活路径
	Auths           string `gorm:"type:varchar(255);comment:权限标识,逗号分隔" json:"auths"`       // 权限标识
	FrameSrc        string `gorm:"type:varchar(255);comment:内嵌 iframe 地址" json:"frameSrc"` // iframe 地址
	FrameLoading    bool   `gorm:"comment:是否显示 iframe 加载动画" json:"frameLoading"`           // 是否加载动画
	KeepAlive       bool   `gorm:"comment:是否缓存组件" json:"keepAlive"`                        // 是否缓存
	HiddenTag       bool   `gorm:"comment:是否隐藏标签" json:"hiddenTag"`                        // 是否隐藏标签
	FixedTag        bool   `gorm:"comment:是否固定标签" json:"fixedTag"`                         // 是否固定标签
	ShowLink        bool   `gorm:"comment:是否显示链接" json:"showLink"`                         // 是否显示链接
	ShowParent      bool   `gorm:"comment:是否显示父级菜单" json:"showParent"`                     // 是否显示父级
}

func (Menu) TableName() string {
	return TablePrefix + "menu"
}

// 获取单个菜单（条件查询）
func GetMenu(where map[string]any) (*Menu, error) {
	var menu Menu
	err := db.Where(where).First(&menu).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &menu, nil
}

// 获取菜单列表（条件+分页）
func GetMenuList(pageNum, pageSize int, where map[string]any) ([]*Menu, error) {
	var menus []*Menu

	query := db.Model(&Menu{})
	if len(where) > 0 {
		query = query.Where(where)
	}

	err := query.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("rank ASC").
		Find(&menus).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return menus, nil
}

// 创建菜单
func CreateMenu(menu Menu) error {
	return db.Create(&menu).Error
}

// 更新菜单（通过ID）
func UpdateMenu(id int, data map[string]any) error {
	return db.Model(&Menu{}).Where("id = ?", id).Updates(data).Error
}

// 删除菜单（单个）
func DeleteMenu(id int) error {
	// 先查询出所有菜单，用于计算要删除的子节点
	var menus []Menu
	if err := db.Find(&menus).Error; err != nil {
		return err
	}

	idsToDelete := []int{id}
	queue := []int{id}

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		for _, m := range menus {
			if m.ParentID == currentID {
				idsToDelete = append(idsToDelete, m.ID)
				queue = append(queue, m.ID)
			}
		}
	}

	return db.Where("id IN ?", idsToDelete).Delete(&Menu{}).Error
}

// 获取所有菜单（不分页）
func GetAllMenus(where map[string]any) ([]*Menu, error) {
	var menus []*Menu
	err := db.Where(where).Order("`rank` ASC").Find(&menus).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return menus, nil
}
