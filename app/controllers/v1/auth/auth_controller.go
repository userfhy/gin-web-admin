package authController

import (
	model "gin-test/app/models"
	"gin-test/app/service/v1/user"
	"gin-test/common"
	"gin-test/utils"
	"gin-test/utils/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary User Login
// @Description 用户登录
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body userService.AuthStruct true "user login"
// @Success 200 {object} common.Response
// @Router /login [post]
func UserLogin(c *gin.Context) {
	appG := common.Gin{C: c}

	// 绑定 payload 到结构体
	var userLogin userService.AuthStruct
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	// 验证绑定结构体参数
	err, parameterErrorStr := common.CheckBindStructParameter(userLogin, c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, parameterErrorStr, nil)
		return
	}

	data := make(map[string]interface{})
	RCode := code.InvalidParams
	isExist, userId, roleKey, isAdmin := model.CheckAuth(userLogin.Username, userLogin.Password)
	if isExist {
		token, err := utils.GenerateToken(utils.Claims{
			UserId:   userId,
			Username: userLogin.Username,
			RoleKey:  roleKey,
			IsAdmin:  isAdmin,
		})
		if err != nil {
			RCode = code.ErrorAuthToken
		} else {
			// 设置登录时间
			userService.SetLoggedTime(userId)
			data["token"] = token

			RCode = code.SUCCESS
		}
	} else {
		RCode = code.ErrorUserPasswordInvalid
	}

	appG.Response(http.StatusOK, RCode, code.GetMsg(RCode), data)
}

// @Summary User Logout
// @Description 用户登出
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Success 200 {object} common.Response
// @Router /user/logout [put]
func UserLogout(c *gin.Context) {
	appG := common.Gin{C: c}
	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	userService.JoinBlockList(user.UserId, c.GetHeader("Authorization")[7:])
	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// @Summary 修改密码
// @Description 密码修改
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body userService.ChangePasswordStruct true "user change password"
// @Success 200 {object} common.Response
// @Router /user/change_password [put]
func ChangePassword(c *gin.Context) {
	appG := common.Gin{C: c}

	var userChangePassword userService.ChangePasswordStruct
	err := c.ShouldBindJSON(&userChangePassword)

	if utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	err, parameterErrorStr := common.CheckBindStructParameter(userChangePassword, c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, parameterErrorStr, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	isExist, userId, _, _ := model.CheckAuth(user.Username, userChangePassword.OldPassword)
	if !isExist {
		RCode := code.ErrorUserOldPasswordInvalid
		appG.Response(http.StatusOK, RCode, code.GetMsg(RCode), nil)
		return
	}

	passwordChangeSuccessful := userService.ChangeUserPassword(userId, userChangePassword.NewPassword)
	if !passwordChangeSuccessful {
		appG.Response(http.StatusOK, code.UnknownError, code.GetMsg(code.UnknownError), nil)
		return
	}

	userService.JoinBlockList(user.UserId, c.GetHeader("Authorization")[7:])
	appG.Response(http.StatusOK, code.SUCCESS, code.GetMsg(code.SUCCESS), nil)
}

// @Summary 当前登录用户信息
// @Description 当前登录用户信息
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Success 200 {object} common.Response
// @Router /user/logged_in [get]
func GetLoggedInUser(c *gin.Context) {
	appG := common.Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	data := make(map[string]interface{}, 0)
	data["user_id"] = user.UserId
	data["user_name"] = user.Username
	data["roles"] = [...]string{user.RoleKey}
	data["permissions"] = [...]string{""}
	if user.IsAdmin {
		data["permissions"] = [...]string{"*:*:*"}
	}
	//data["permissions"] = [...]string{"system:sysmenu:add"}
	appG.Response(http.StatusOK, code.SUCCESS, "当前登录用户信息", data)
}
