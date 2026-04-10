package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant"
	ut "github.com/go-playground/universal-translator"

	//"gopkg.in/go-playground/validator.v9"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitValidate() {
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()
}

func CheckBindStructParameter(s any, c *gin.Context) (string, error) {
	v, _ := c.Get("trans")

	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = Uni.GetTranslator("zh")
	}

	if err := Validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return strings.Join(sliceErrs, ","), errs
	}

	return "", nil
}
