package utils

import (
    "gin-test/common"
    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
    "log"
)

type Page struct {
    P uint `json:"p" validate:"required,min=1"`
    N uint `json:"n" validate:"required,min=1"`
}

// GetPage get page parameters
func GetPage(c *gin.Context) (int, int) {
    currentPage := 0

    v, _ := c.Get("trans")

    trans, ok := v.(ut.Translator)
    if !ok {
        trans, _ = common.Uni.GetTranslator("zh")
    }

    var p Page
    if err := c.ShouldBindQuery(&p); err != nil {
       log.Println(err)
    }
    
    err := common.Validate.Struct(p)
    if err != nil {
       errs := err.(validator.ValidationErrors)
       var sliceErrs [] string
       for _, e := range errs {
           sliceErrs = append(sliceErrs, e.Translate(trans))
       }
       log.Println(sliceErrs)
    }

    page := c.DefaultQuery("p", "0")
    limit := c.DefaultQuery("limit", "15")

    log.Println(page)
    log.Println(limit)
    
    //if page > 0 {
    //    currentPage = (page - 1) * limit
    //}
    
    return currentPage, 0
}