package casbinService

import (
	model "gin-web-admin/app/models"
	"log"
)

type CasbinStruct struct {
	PageNum  int
	PageSize int
	V0       string // role
	V1       string // path
	V2       string // method
}

type AddCasbinStruct struct {
	V0 string `json:"v0" form:"v0" validate:"required,min=3,max=20" minLength:"3",maxLength:"20"` // role_key
	V1 string `json:"v1" form:"v1" validate:"required,min=4,max=30" minLength:"4",maxLength:"30"` // path
	V2 string `json:"v2" form:"v2" validate:"required,min=2,max=8" minLength:"2",maxLength:"8"`   // method
}

func CreateCasbin(n AddCasbinStruct) error {
	return model.CreatCasbin(model.CasbinRuleM{
		Ptype: "p",
		V0:    n.V0,
		V1:    n.V1,
		V2:    n.V2,
	})
}

func UpdateCasbin(id int, u AddCasbinStruct) bool {
	wheres := make(map[string]interface{})
	wheres["id"] = id

	updates := make(map[string]interface{})
	updates["v1"] = u.V1
	updates["v2"] = u.V2
	error, rowsAffected := model.Update(&model.CasbinRuleM{}, wheres, updates)
	if rowsAffected == 0 {
		log.Println("修改Casbin失败！")
		log.Println(error)
		return false
	}
	return true
}

func (c *CasbinStruct) getConditionMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	//maps["deleted_at"] = nil
	if c.V0 != "" {
		maps["v0"] = c.V0
	}

	if c.V1 != "" {
		maps["v1 like"] = "%" + c.V1 + "%"
	}

	if c.V2 != "" {
		maps["v2"] = c.V2
	}
	// log.Println(c)
	// log.Println(maps)
	return maps
}

func (c *CasbinStruct) Count() (int64, error) {
	return model.GetTotal(model.CasbinRuleM{}, c.getConditionMaps())
}

func (c *CasbinStruct) GetAll() ([]*model.CasbinRuleM, error) {
	casbins, err := model.GetCasbinRuleList(c.PageNum, c.PageSize, c.getConditionMaps())
	if err != nil {
		return nil, err
	}

	return casbins, nil
}
