package authController

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authService "gin-web-admin/app/service/v1/auth"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/common"
	"gin-web-admin/utils"
	"gin-web-admin/utils/code"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitValidate()
}

type mockAuthService struct {
	loginFunc           func(userService.AuthStruct, string) (authService.LoginResult, error)
	refreshFunc         func(string, string, string) (map[string]any, error)
	changePasswordFunc  func(string, userService.ChangePasswordStruct) error
	buildLoggedUserFunc func(*utils.Claims) map[string]any
}

func (m *mockAuthService) Login(payload userService.AuthStruct, ip string) (authService.LoginResult, error) {
	return m.loginFunc(payload, ip)
}

func (m *mockAuthService) RefreshAccessToken(token, clientIP, userAgent string) (map[string]any, error) {
	if m.refreshFunc != nil {
		return m.refreshFunc(token, clientIP, userAgent)
	}
	return map[string]any{}, nil
}

func (m *mockAuthService) Logout(uint, string) {}

func (m *mockAuthService) ChangePassword(username string, payload userService.ChangePasswordStruct) error {
	if m.changePasswordFunc != nil {
		return m.changePasswordFunc(username, payload)
	}
	return nil
}

func (m *mockAuthService) BuildLoggedInUserData(claims *utils.Claims) map[string]any {
	if m.buildLoggedUserFunc != nil {
		return m.buildLoggedUserFunc(claims)
	}
	return map[string]any{}
}

func TestHandler_UserLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAuthService{
		loginFunc: func(userService.AuthStruct, string) (authService.LoginResult, error) {
			return authService.LoginResult{
				UserID:       1,
				Username:     "admin",
				RoleKey:      "admin",
				IsAdmin:      true,
				AccessToken:  "token",
				RefreshToken: "refresh",
			}, nil
		},
	}
	handler := NewHandler(mockSvc)
	router := gin.New()
	router.POST("/login", handler.UserLogin)

	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "pass",
	})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]any
	_ = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, float64(code.SUCCESS), response["code"])
	assert.NotNil(t, response["data"])
}

func TestHandler_UserLoginInvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAuthService{
		loginFunc: func(userService.AuthStruct, string) (authService.LoginResult, error) {
			return authService.LoginResult{}, authService.ErrInvalidCredentials
		},
	}
	handler := NewHandler(mockSvc)
	router := gin.New()
	router.POST("/login", handler.UserLogin)

	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "wrong",
	})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]any
	_ = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, float64(code.ErrorUserPasswordInvalid), response["code"])
}
