package utils

import (
	"gin-web-admin/common"
	"gin-web-admin/utils/com"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	//Page     int         `json:"page"`
	PageSize int `json:"page_size"`
}

type Page struct {
	P uint `json:"p" form:"p" validate:"required,numeric,min=1"`
	N uint `json:"n" form:"n" validate:"required,numeric,min=1"`
}

// GetPage get page parameters
func GetPage(c *gin.Context) (error, string, int, int) {
	currentPage := 0

	// 绑定 query 参数到结构体
	var p Page
	if err := c.ShouldBindQuery(&p); err != nil {
		return err, "参数绑定失败,请检查传递参数类型！", 0, 0
	}

	// 验证绑定结构体参数
	err, parameterErrorStr := common.CheckBindStructParameter(p, c)
	if err != nil {
		return err, parameterErrorStr, 0, 0
	}

	page := com.StrTo(c.DefaultQuery("p", "0")).MustInt()
	limit := com.StrTo(c.DefaultQuery("n", "15")).MustInt()

	if page > 0 {
		currentPage = (page - 1) * limit
	}

	return nil, "", currentPage, limit
}
