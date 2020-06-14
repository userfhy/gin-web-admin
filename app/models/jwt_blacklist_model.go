package model

type JwtBlacklist struct {
	BaseModel
	UserID uint   `json:"user_id"`
	Jwt    string `gorm:"type:text"`
}

func (JwtBlacklist) TableName() string {
	return TablePrefix + "jwt_blacklist"
}

func CreateBlockList(userId uint, jwt string) error {
	table := JwtBlacklist{UserID: userId, Jwt: jwt}
	db.NewRecord(table)
	res := db.Create(&table)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
