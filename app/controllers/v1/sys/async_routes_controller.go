package sysController

import (
	"net/http"

	"gin-web-admin/utils"

	"github.com/gin-gonic/gin"
)

// GetAsyncRoutes returns async routes generated from gin_menu.
// It filters menus by current user's role via gin_role_menu.
// GET /get-async-routes
func (h *Handler) GetAsyncRoutes(c *gin.Context) {
	claimsAny, _ := c.Get("claims")
	claims, _ := claimsAny.(*utils.Claims)
	roleKey := ""
	if claims != nil {
		roleKey = claims.RoleKey
	}

	routes, err := h.service.GetAsyncRoutes(roleKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    routes,
	})
}
