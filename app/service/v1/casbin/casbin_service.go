package casbinService

import (
    model "gin-test/app/models"
    "log"
)

type CasbinStruct struct {
    PageNum  int
    PageSize int
}

type AddCasbinStruct struct {
    V0 string `json:"v0" form:"v0" validate:"required,min=3,max=20" minLength:"3",maxLength:"20"` // role_key
    V1 string `json:"v1" form:"v1" validate:"required,min=4,max=30" minLength:"4",maxLength:"30"` // path
    V2 string `json:"v2" form:"v2" validate:"required,min=2,max=8" minLength:"2",maxLength:"8"` // method
}

func CreateCasbin(n AddCasbinStruct) error{
    return model.CreatCasbin(model.CasbinRule{
        PType: "p",
        V0: n.V0,
        V1: n.V1,
        V2: n.V2,
    })
}

func UpdateCasbin(id int, u AddCasbinStruct) bool{
    wheres := make(map[string]interface{})
    wheres["id"] = id

    updates := make(map[string]interface{})
    updates["v1"] = u.V1
    updates["v2"] = u.V2
    error, rowsAffected := model.Update(&model.CasbinRule{}, wheres, updates)
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
    return maps
}

func (c *CasbinStruct) Count() (int, error) {
    return model.GetTotal(model.CasbinRule{}, c.getConditionMaps())
}

func (c *CasbinStruct) GetAll() ([]*model.CasbinRule, error) {
    casbins, err := model.GetCasbinRuleList(c.PageNum, c.PageSize, c.getConditionMaps())
    if err != nil {
        return nil, err
    }

    return casbins, nil
}