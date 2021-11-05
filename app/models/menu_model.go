package model

type Menu struct {
	MenuId     int    `json:"menu_id" gorm:"primary_key;AUTO_INCREMENT"`
	ParentId   int    `json:"parent_id" gorm:"type:int(11);"`
	Sort       int    `json:"sort" gorm:"type:int(4);"`
	MenuName   string `json:"menu_name" gorm:"type:varchar(11);comment:'路由名称'"`
	Path       string `json:"path" gorm:"type:varchar(128);comment:'路由路径'"`
	Paths      string `json:"paths" gorm:"type:varchar(128);"`
	Component  string `json:"component" gorm:"type:varchar(255);comment:'组件路径'"`
	Title      string `json:"title" gorm:"type:varchar(64);comment:'菜单标题'"`
	Icon       string `json:"icon" gorm:"type:varchar(128);"`
	MenuType   string `json:"menu_type" gorm:"type:varchar(1);"` //"M"：目录 "C"：菜单 "F"：按钮
	Permission string `json:"permission" gorm:"type:varchar(32);"`
	Visible    string `json:"visible" gorm:"type:int(1);DEFAULT:0;"`
	IsFrame    string `json:"is_frame" gorm:"type:int(1);DEFAULT:0;"` // 是否是外链
	Params     string `json:"params" gorm:"-"`
	RoleId     int    `gorm:"-"`
	Children   []Menu `json:"children" gorm:"-"`
	IsSelect   bool   `json:"is_select" gorm:"-"`
	BaseModelNoId
}

func (Menu) TableName() string {
	return TablePrefix + "menu"
}
