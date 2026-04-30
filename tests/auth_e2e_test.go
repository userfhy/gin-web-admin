package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authController "gin-web-admin/app/controllers/v1/auth"
	authService "gin-web-admin/app/service/v1/auth"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type e2eStubAuthService struct {
	result authService.LoginResult
	err    error
}

func (s *e2eStubAuthService) Login(_ userService.AuthStruct, _, _ string) (authService.LoginResult, error) {
	return s.result, s.err
}

func (s *e2eStubAuthService) RefreshAccessToken(string, string, string) (map[string]any, error) {
	return map[string]any{}, nil
}

func (s *e2eStubAuthService) Logout(uint, string) {}

func (s *e2eStubAuthService) ChangePassword(string, userService.ChangePasswordStruct) error {
	return nil
}

func (s *e2eStubAuthService) BuildLoggedInUserData(*utils.Claims) map[string]any {
	return map[string]any{}
}

func TestAuthLoginEndpoint(t *testing.T) {
	common.InitValidate()
	gin.SetMode(gin.TestMode)

	router := gin.New()
	result := authService.LoginResult{
		UserID:       3,
		Username:     "demo",
		RoleKey:      "editor",
		IsAdmin:      false,
		AccessToken:  "access",
		RefreshToken: "refresh",
	}
	handler := authController.NewHandler(&e2eStubAuthService{result: result})
	router.POST("/v1/api/login", handler.UserLogin)

	body, _ := json.Marshal(map[string]string{
		"username": "demo",
		"password": "demo-pass",
	})
	req := httptest.NewRequest(http.MethodPost, "/v1/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var payload struct {
		Code int                    `json:"code"`
		Data map[string]interface{} `json:"data"`
	}
	_ = json.Unmarshal(resp.Body.Bytes(), &payload)
	assert.Equal(t, code.SUCCESS, payload.Code)
	assert.Equal(t, "demo", payload.Data["username"])
	assert.Equal(t, "access", payload.Data["accessToken"])
}
