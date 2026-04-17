package middleware

import (
	"gin-web-admin/common"
	"gin-web-admin/utils/logging"
	"sync"

	"github.com/gin-gonic/gin"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var (
	registerZhOnce   sync.Once
	registerEnOnce   sync.Once
	registerZhTwOnce sync.Once
)

// 设置Translation
func TranslationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := common.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			registerZhOnce.Do(func() {
				if err := zh_translations.RegisterDefaultTranslations(common.Validate, trans); err != nil {
					logging.Warnf("register zh translations failed: %v", err)
				}
			})
		case "en":
			registerEnOnce.Do(func() {
				if err := en_translations.RegisterDefaultTranslations(common.Validate, trans); err != nil {
					logging.Warnf("register en translations failed: %v", err)
				}
			})
		case "zh_tw":
			registerZhTwOnce.Do(func() {
				if err := zh_tw_translations.RegisterDefaultTranslations(common.Validate, trans); err != nil {
					logging.Warnf("register zh_tw translations failed: %v", err)
				}
			})
		default:
			registerZhOnce.Do(func() {
				if err := zh_translations.RegisterDefaultTranslations(common.Validate, trans); err != nil {
					logging.Warnf("register default translations failed: %v", err)
				}
			})
		}

		//设置trans到context
		c.Set("trans", trans)
		c.Next()
	}
}
