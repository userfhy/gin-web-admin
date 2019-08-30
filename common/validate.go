package common

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/locales/en"
    "github.com/go-playground/locales/zh"
    "github.com/go-playground/locales/zh_Hant_TW"
    "github.com/go-playground/universal-translator"
    "gopkg.in/go-playground/validator.v9"
    "strings"
)

var (
    Uni      *ut.UniversalTranslator
    Validate *validator.Validate
)

func InitValidate()  {
    en := en.New()
    zh := zh.New()
    zh_tw := zh_Hant_TW.New()
    Uni = ut.New(en, zh, zh_tw)
    Validate = validator.New()
}

func CheckBindStructParameter(s interface{}, c *gin.Context) (error, string) {
    v, _ := c.Get("trans")

    trans, ok := v.(ut.Translator)
    if !ok {
        trans, _ = Uni.GetTranslator("zh")
    }

    err := Validate.Struct(s)
    if err != nil {
        errs := err.(validator.ValidationErrors)
        var sliceErrs [] string
        for _, e := range errs {
            sliceErrs = append(sliceErrs, e.Translate(trans))
        }
        //log.Println(errs.Translate(trans))
        return errs, strings.Join(sliceErrs, ",")
    }

    return nil, ""
}