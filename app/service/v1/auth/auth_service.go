package authService

import (
	"errors"
	"time"

	model "gin-web-admin/app/models"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserDisabled       = errors.New("user disabled")
)

type Service struct {
	store       *data.Store
	userService *userService.Service
}

type LoginResult struct {
	UserID       uint
	Username     string
	RoleKey      string
	IsAdmin      bool
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func NewService(store *data.Store, userSvc *userService.Service) *Service {
	if userSvc == nil {
		panic("auth service requires user service dependency")
	}
	return &Service{store: store, userService: userSvc}
}

func (s *Service) Login(payload userService.AuthStruct) (LoginResult, error) {
	isExist, userID, roleKey, isAdmin, status := model.CheckAuth(payload.Username, payload.Password)
	if !isExist {
		return LoginResult{}, ErrInvalidCredentials
	}
	if status == 0 {
		return LoginResult{}, ErrUserDisabled
	}

	claims := utils.Claims{
		UserId:   userID,
		Username: payload.Username,
		RoleKey:  roleKey,
		IsAdmin:  isAdmin,
	}

	accessToken, expiresAt, err := utils.GenerateToken(claims)
	if err != nil {
		return LoginResult{}, err
	}
	refreshToken, _, err := utils.GenerateRefreshToken(claims)
	if err != nil {
		return LoginResult{}, err
	}

	oldRefresh, err := s.userService.SetLoggedUserInfo(userID, refreshToken)
	if err != nil {
		return LoginResult{}, err
	}
	if oldRefresh != "" && oldRefresh != refreshToken {
		s.userService.JoinBlockList(userID, oldRefresh)
	}

	return LoginResult{
		UserID:       userID,
		Username:     payload.Username,
		RoleKey:      roleKey,
		IsAdmin:      isAdmin,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

func (s *Service) RefreshAccessToken(refreshToken string) (map[string]any, error) {
	return s.userService.RefreshAccessToken(refreshToken)
}

func (s *Service) Logout(userID uint, token string) {
	if token == "" {
		return
	}
	s.userService.JoinBlockList(userID, token)
}

func (s *Service) ChangePassword(username string, payload userService.ChangePasswordStruct) error {
	isExist, userID, _, _, status := model.CheckAuth(username, payload.OldPassword)
	if !isExist {
		return ErrInvalidCredentials
	}
	if status == 0 {
		return ErrUserDisabled
	}
	ok := s.userService.ChangeUserPassword(userID, payload.NewPassword)
	if !ok {
		return errors.New("change password failed")
	}
	return nil
}

func (s *Service) BuildLoggedInUserData(claims *utils.Claims) map[string]any {
	data := map[string]any{
		"user_id":     claims.UserId,
		"username":    claims.Username,
		"roles":       [...]string{claims.RoleKey},
		"permissions": [...]string{""},
	}
	if claims.IsAdmin {
		data["permissions"] = [...]string{"*:*:*"}
	}
	return data
}
