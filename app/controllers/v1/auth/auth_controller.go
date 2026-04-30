package authController

import (
	"errors"
	"net/http"

	authService "gin-web-admin/app/service/v1/auth"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/logging"

	"github.com/gin-gonic/gin"
)

var ExpireTimeFormat = "2006/01/02 15:04:05"

type AuthService interface {
	Login(payload userService.AuthStruct, clientIP, userAgent string) (authService.LoginResult, error)
	RefreshAccessToken(refreshToken, clientIP, userAgent string) (map[string]any, error)
	Logout(userID uint, token string)
	ChangePassword(username string, payload userService.ChangePasswordStruct) error
	BuildLoggedInUserData(claims *utils.Claims) map[string]any
}

type Handler struct {
	service AuthService
}

func NewHandler(service AuthService) *Handler {
	if service == nil {
		panic("auth handler requires non-nil service")
	}
	return &Handler{service: service}
}

// @Summary		User Login
// @Description	用户登录
// @Accept			json
// @Produce		json
// @Tags			Auth
// @Param			payload	body		userService.AuthStruct	true	"user login"
// @Success		200		{object}	common.Response
// @Router			/login [post]
func (h *Handler) UserLogin(c *gin.Context) {
	appG := common.Gin{C: c}

	// 绑定 payload 到结构体
	var userLogin userService.AuthStruct
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	// 验证绑定结构体参数
	parameterErrorStr, err := common.CheckBindStructParameter(userLogin, c)
	if utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
		return
	}

	clientIP := c.ClientIP()
	result, err := h.service.Login(userLogin, clientIP, c.GetHeader("User-Agent"))
	if err != nil {
		switch {
		case errors.Is(err, authService.ErrInvalidCredentials):
			appG.Response(http.StatusOK, code.ErrorUserPasswordInvalid, code.GetMsg(code.ErrorUserPasswordInvalid), nil)
		case errors.Is(err, authService.ErrUserDisabled):
			appG.Response(http.StatusOK, code.ErrorAuth, "该用户已被禁用", nil)
		case errors.Is(err, authService.ErrAccountLocked):
			data := gin.H{}
			message := "账号已锁定，请稍后再试"
			if lockedUntil, ok := authService.LockedUntilFromError(err); ok {
				formatted := lockedUntil.Format(ExpireTimeFormat)
				message = "账号已锁定，解锁时间：" + formatted
				data["lockedUntil"] = formatted
			}
			appG.Response(http.StatusTooManyRequests, code.ErrorAuth, message, data)
		case errors.Is(err, authService.ErrIPNotAllowed):
			appG.Response(http.StatusForbidden, code.ErrorAuth, "当前 IP 未被授权登录", nil)
		default:
			utils.HandleError(c, http.StatusInternalServerError, code.ERROR, "登录失败", err)
		}
		return
	}

	data := map[string]any{
		"accessToken":  result.AccessToken,
		"token":        result.AccessToken,
		"refreshToken": result.RefreshToken,
		"username":     result.Username,
		"nickname":     result.Username,
		"roles":        [1]string{result.RoleKey},
		"expires":      result.ExpiresAt.Format(ExpireTimeFormat),
	}

	appG.Response(http.StatusOK, code.SUCCESS, "用户登录成功", data)
}

// @Summary		Auth RefreshAccessToken
// @Description	刷新用户access_token
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			Auth
// @Param			payload	body		userService.RefreshAccessTokenStruct	true	"根据refresh_token 刷新access_token"
// @Success		200		{object}	common.Response
// @Router			/refresh_token [post]
func (h *Handler) RefreshAccessToken(c *gin.Context) {
	appG := common.Gin{C: c}
	var refreshAccessTokenhStruct userService.RefreshAccessTokenStruct
	if err := c.ShouldBindJSON(&refreshAccessTokenhStruct); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	data, err := h.service.RefreshAccessToken(refreshAccessTokenhStruct.RefreshToken, c.ClientIP(), c.GetHeader("User-Agent"))
	if utils.HandleError(c, http.StatusOK, code.ErrorAuthToken, "access_token刷新失败", err) {
		logging.Println("Error token: ", err)
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "刷新access_token成功！", data)
}

// @Summary		User Logout
// @Description	用户登出
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			User
// @Success		200	{object}	common.Response
// @Router			/user/logout [post]
func (h *Handler) UserLogout(c *gin.Context) {
	appG := common.Gin{C: c}
	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	token := c.GetHeader("Authorization")
	if len(token) > 7 {
		h.service.Logout(user.UserId, token[7:])
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", nil)
}

// @Summary		修改密码
// @Description	密码修改
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			User
// @Param			payload	body		userService.ChangePasswordStruct	true	"user change password"
// @Success		200		{object}	common.Response
// @Router			/user/change_password [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	appG := common.Gin{C: c}

	var userChangePassword userService.ChangePasswordStruct
	if err := c.ShouldBindJSON(&userChangePassword); utils.HandleError(c, http.StatusBadRequest, http.StatusBadRequest, "参数绑定失败", err) {
		return
	}

	parameterErrorStr, err := common.CheckBindStructParameter(userChangePassword, c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, parameterErrorStr, nil)
		return
	}

	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	if err := h.service.ChangePassword(user.Username, userChangePassword); err != nil {
		switch {
		case errors.Is(err, authService.ErrInvalidCredentials):
			appG.Response(http.StatusOK, code.ErrorUserOldPasswordInvalid, code.GetMsg(code.ErrorUserOldPasswordInvalid), nil)
		case errors.Is(err, authService.ErrUserDisabled):
			appG.Response(http.StatusOK, code.ErrorAuth, "该用户已被禁用", nil)
		case errors.Is(err, authService.ErrAccountLocked):
			appG.Response(http.StatusTooManyRequests, code.ErrorAuth, "账号已锁定，请稍后再试", nil)
		default:
			appG.Response(http.StatusOK, code.UnknownError, err.Error(), nil)
		}
		return
	}

	token := c.GetHeader("Authorization")
	if len(token) > 7 {
		h.service.Logout(user.UserId, token[7:])
	}
	appG.Response(http.StatusOK, code.SUCCESS, code.GetMsg(code.SUCCESS), nil)
}

// @Summary		当前登录用户信息
// @Description	当前登录用户信息
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			User
// @Success		200	{object}	common.Response
// @Router			/user/logged_in [get]
func (h *Handler) GetLoggedInUser(c *gin.Context) {
	appG := common.Gin{C: c}

	claims, _ := c.Get("claims")
	user := claims.(*utils.Claims)

	data := h.service.BuildLoggedInUserData(user)
	appG.Response(http.StatusOK, code.SUCCESS, "当前登录用户信息", data)
}
