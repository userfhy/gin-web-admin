package casbin

import model "gin-test/app/models"

type CasbinStruct struct {
    PageNum  int
    PageSize int
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