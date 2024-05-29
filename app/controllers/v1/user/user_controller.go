package userController

import (
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 创建用户
// @Description 创建新用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body userService.AddUserStruct true "create new user"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user [post]
func CreateUser(c *gin.Context) {
	appG := common.Gin{C: c}

	var newUser userService.AddUserStruct
	err := c.ShouldBindJSON(&newUser)
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	err, parameterErrorStr := common.CheckBindStructParameter(newUser, c)
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
		return
	}

	err = userService.CreateUser(newUser)
	if utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "添加新用户失败！", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "用户添加成功", nil)
}

// @Summary 用户列表
// @Description 获取用户列表
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param p query int true "page number"
// @Param n query int true "page limit"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user [get]
func GetUsers(c *gin.Context) {
	appG := common.Gin{C: c}

	_, errStr, p, n := utils.GetPage(c)
	if errStr != "" {
		appG.Response(http.StatusBadRequest, code.InvalidParams, errStr, nil)
		return
	}

	var userServiceObj userService.UserStruct
	userServiceObj.PageNum = p
	userServiceObj.PageSize = n

	err := c.ShouldBindQuery(&userServiceObj)
	if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	total, err := userServiceObj.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "获取页数失败"+err.Error(), nil)
		return
	}

	userArr, err := userServiceObj.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, "服务器错误"+err.Error(), nil)
		return
	}

	/*    var uIndex []userService.TestList
	      for i := range userArr {
	          uIndex = append(uIndex, userService.TestList{Index: i + 1, Auth: userArr[i]})
	          //log.Println(userArr[i])
	      }
	      log.Println(uIndex)*/

	data := utils.PageResult{
		List:        userArr,
		Total:       total,
		CurrentPage: p,
		PageSize:    n,
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", data)
}
