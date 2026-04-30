package userController

import (
	"net/http"
	"strconv"

	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"
	"gin-web-admin/utils/query"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *userService.Service
}

func NewHandler(service *userService.Service) *Handler {
	if service == nil {
		panic("user handler requires non-nil service")
	}
	return &Handler{service: service}
}

// @Summary		创建用户
// @Description	创建新用户
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			User
// @Param			payload	body		userService.AddUserStruct	true	"create new user"
// @Success		200		{object}	common.Response
// @Failure		500		{object}	common.Response
// @Router			/user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	appG := common.Gin{C: c}

	var newUser userService.AddUserStruct
	if err := c.ShouldBindJSON(&newUser); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}

	if parameterErrorStr, err := common.CheckBindStructParameter(newUser, c); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
		return
	}

	if err := h.service.CreateUser(newUser); utils.HandleError(c, http.StatusInternalServerError, http.StatusInternalServerError, "添加新用户失败！", err) {
		return
	}

	appG.Response(http.StatusOK, code.SUCCESS, "用户添加成功", nil)
}

// @Summary		用户列表
// @Description	获取用户列表
// @Accept			json
// @Produce		json
// @Security		ApiKeyAuth
// @Tags			User
// @Param			p	query		int	true	"page number"
// @Param			n	query		int	true	"page limit"
// @Success		200	{object}	common.Response
// @Failure		500	{object}	common.Response
// @Router			/user [get]
func (h *Handler) GetUsers(c *gin.Context) {
	appG := common.Gin{C: c}

	pg, err := utils.GetPagination(c)
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	var userServiceObj userService.UserStruct
	userServiceObj.Pagination = pg.Clone()

	if err := c.ShouldBindQuery(&userServiceObj); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "参数绑定失败", nil)
		return
	}

	filterBuilder := query.NewBuilder().IsNull("deleted_at")
	if err := filterBuilder.FromQuery(c, query.RuleSet{
		"username": query.Rule{Field: "username", Op: query.OpLike},
		"status":   query.Rule{Field: "status", Op: query.OpEqual, Parser: query.IntParser()},
	}); err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}
	userServiceObj.Conditions = filterBuilder.Build()

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

	appG.Response(http.StatusOK, code.SUCCESS, "ok", pg.Result(userArr, total))
}

func (h *Handler) ResetPassword(c *gin.Context) {
	appG := common.Gin{C: c}
	userID, ok := parseUserIDParam(c, appG)
	if !ok {
		return
	}

	var payload userService.ResetPasswordStruct
	if err := c.ShouldBindJSON(&payload); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, "参数绑定失败", err) {
		return
	}
	if parameterErrorStr, err := common.CheckBindStructParameter(payload, c); utils.HandleError(c, http.StatusBadRequest, code.InvalidParams, parameterErrorStr, err) {
		return
	}

	operator := ""
	if claimsValue, exists := c.Get("claims"); exists {
		if claims, ok := claimsValue.(*utils.Claims); ok && claims != nil {
			operator = claims.Username
		}
	}

	if err := h.service.ResetUserPassword(userID, payload.NewPassword, operator, c.ClientIP()); err != nil {
		appG.Response(http.StatusOK, code.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "重置密码成功，用户全部会话已强制下线", nil)
}

func (h *Handler) UnlockUser(c *gin.Context) {
	appG := common.Gin{C: c}
	userID, ok := parseUserIDParam(c, appG)
	if !ok {
		return
	}

	operator := ""
	if claimsValue, exists := c.Get("claims"); exists {
		if claims, ok := claimsValue.(*utils.Claims); ok && claims != nil {
			operator = claims.Username
		}
	}

	if err := h.service.UnlockUser(userID, operator, c.ClientIP()); err != nil {
		appG.Response(http.StatusOK, code.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "解锁成功", nil)
}

func (h *Handler) GetSecurityTimeline(c *gin.Context) {
	appG := common.Gin{C: c}
	userID, ok := parseUserIDParam(c, appG)
	if !ok {
		return
	}

	pg, err := utils.GetPagination(c, utils.WithDefaultPageSize(5), utils.WithMaxPageSize(50))
	if err != nil {
		appG.Response(http.StatusBadRequest, code.InvalidParams, err.Error(), nil)
		return
	}

	events, err := h.service.GetUserSecurityTimeline(userID, pg)
	if err != nil {
		appG.Response(http.StatusInternalServerError, code.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, code.SUCCESS, "ok", events)
}

func parseUserIDParam(c *gin.Context, appG common.Gin) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		appG.Response(http.StatusBadRequest, code.InvalidParams, "用户ID错误", nil)
		return 0, false
	}
	return uint(id), true
}
