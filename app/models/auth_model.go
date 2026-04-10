package model

import (
	"gin-web-admin/utils"

	"gorm.io/gorm"
)

type Auth struct {
	BaseModel
	RoleId           uint      `gorm:"DEFAULT:0;NOT NULL;" json:"role_id"`
	Status           int       `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	LoggedInAt       JSONTime  `json:"logged_in_at"`
	LockedUntil      *JSONTime `json:"locked_until"`
	FailedLoginCount int       `gorm:"type:int;default:0" json:"failed_login_count"`
	LastLoginIP      string    `gorm:"size:64" json:"last_login_ip"`
	Username         string    `gorm:"Size:20;unique;NOT NULL;" json:"user_name"`
	Nickname         string    `gorm:"Size:30;" json:"nickname"`
	Phone            string    `gorm:"Size:30;" json:"phone"`
	Email            string    `gorm:"Size:40;" json:"email"`
	Sex              int       `gorm:"type:tinyint(1) unsigned;DEFAULT:0;NOT NULL;comment:1-女 2-男" json:"sex"`
	Password         string    `gorm:"Size:50;NOT NULL;" json:"-"`
	RefreshToken     string    `gorm:"Size:600;unique;default:''" json:"refresh_token"`
	RoleName         string    `gorm:"-" json:"role_name"`
	Role             Role      `gorm:"foreignkey:RoleId" json:"-"`
}

func (Auth) TableName() string {
	return TablePrefix + "auth"
}

func GetAuthByUsername(username string) (*Auth, error) {
	var auth Auth
	if err := db.Where("username = ?", username).Preload("Role").First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &auth, nil
}

func CreatUser(auth Auth) error {
	res := db.Create(&auth)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func GetUser(maps map[string]any) (*Auth, error) {
	var user *Auth
	err := db.Select("*").Where(maps).Preload(
		"Role", func(db *gorm.DB) *gorm.DB {
			return db.Select("*")
		}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}

// GetTestUsers gets a list of users based on paging constraints
func GetUsers(pg utils.Pagination, where map[string]any) ([]*Auth, error) {
	var user []*Auth

	db, _ := BuildCondition(db, where)
	err := db.Select("*").Scopes(pg.Scope()).Preload(
		"Role", func(db *gorm.DB) *gorm.DB {
			return db.Select("role_id,role_name")
		}).Find(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}
