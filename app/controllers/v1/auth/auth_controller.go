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
    isExist, userId := model.CheckAuth(userLogin.Username, userLogin.Password)
    if isExist {
        token, err := utils.GenerateToken(userLogin.Username, userLogin.Password, userId)
        if err != nil {
            RCode = code.ErrorAuthToken
        } else {
            data["token"] = token

            RCode = code.SUCCESS
        }
    } else {
        RCode = code.ErrorUserPasswordInvalid
    }

    appG.Response(http.StatusOK, RCode, code.GetMsg(RCode), data)
}

// @Summary 当前登录用户信息
// @Description 当前登录用户信息
// @Accept json
// @Produce json
// @Tags Auth
// @Success 200 {object} common.Response
// @Router /user/logged_in [get]
func GetLoggedInUser(c *gin.Context) {
    appG := common.Gin{C: c}

    claims, _ := c.Get("claims")
    user := claims.(*utils.Claims)

    data := make(map[string]interface{}, 0)
    data["user_name"] = user.Username
    appG.Response(http.StatusOK, code.SUCCESS, "当前登录用户信息", data)
}