package authController

import (
    model "gin-test/app/models"
    "gin-test/common"
    "gin-test/utils"
    "gin-test/utils/code"
    "github.com/gin-gonic/gin"
    "net/http"
)


type AuthStruct struct {
    Username string `json:"user_name" form:"user_name" validate:"required,min=1,max=20" minLength:"1",maxLength:"20"`
    Password string `json:"password" form:"password" validate:"required,min=4,max=20" minLength:"4",maxLength:"20"`
}

// @Summary User Login
// @Description Test db Connections.
// @Accept json
// @Produce json
// @Tags Auth
// @Param payload body authController.AuthStruct true "user login"
// @Success 200 {object} common.Response
// @Router /login [post]
func GetAuth(c *gin.Context) {
    appG := common.Gin{C: c}

    // 绑定 payload 到结构体
    var userLogin AuthStruct
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