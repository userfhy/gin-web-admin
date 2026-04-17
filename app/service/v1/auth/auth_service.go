package authService

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	model "gin-web-admin/app/models"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/security"
	"gin-web-admin/utils/setting"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserDisabled       = errors.New("user disabled")
	ErrAccountLocked      = errors.New("account locked")
	ErrIPNotAllowed       = errors.New("ip not allowed")
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

func (s *Service) Login(payload userService.AuthStruct, clientIP string) (LoginResult, error) {
	if !security.IsIPWhitelisted(clientIP) {
		s.logLogin(payload.Username, 0, clientIP, false, "IP not allowed")
		return LoginResult{}, ErrIPNotAllowed
	}

	user, err := model.GetAuthByUsername(payload.Username)
	if err != nil {
		return LoginResult{}, err
	}
	if user == nil {
		s.logLogin(payload.Username, 0, clientIP, false, "invalid credentials")
		return LoginResult{}, ErrInvalidCredentials
	}

	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		until := user.LockedUntil.Format(time.RFC3339)
		s.logLogin(user.Username, user.ID, clientIP, false, "account locked until "+until)
		return LoginResult{}, ErrAccountLocked
	}
	if user.Status == 0 {
		s.logLogin(user.Username, user.ID, clientIP, false, "user disabled")
		return LoginResult{}, ErrUserDisabled
	}
	if user.Password != utils.EncodeUserPassword(payload.Password) {
		s.handleLoginFailure(user, clientIP, "invalid password")
		return LoginResult{}, ErrInvalidCredentials
	}

	claims := utils.Claims{
		UserId:   user.ID,
		Username: user.Username,
		RoleKey:  user.Role.RoleKey,
		IsAdmin:  user.Role.IsAdmin,
	}

	accessToken, expiresAt, err := utils.GenerateToken(claims)
	if err != nil {
		return LoginResult{}, err
	}
	refreshToken, _, err := utils.GenerateRefreshToken(claims)
	if err != nil {
		return LoginResult{}, err
	}

	oldRefresh, err := s.userService.SetLoggedUserInfo(user.ID, refreshToken, clientIP)
	if err != nil {
		return LoginResult{}, err
	}
	if oldRefresh != "" && oldRefresh != refreshToken {
		s.userService.JoinBlockList(user.ID, oldRefresh)
	}
	if err := s.userService.SaveOnlineSession(accessToken, &claims, clientIP, ""); err != nil {
		return LoginResult{}, err
	}

	s.logLogin(user.Username, user.ID, clientIP, true, "login success")

	return LoginResult{
		UserID:       user.ID,
		Username:     user.Username,
		RoleKey:      user.Role.RoleKey,
		IsAdmin:      user.Role.IsAdmin,
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
	s.userService.RemoveOnlineSession(userID)
}

func (s *Service) ChangePassword(username string, payload userService.ChangePasswordStruct) error {
	user, err := model.GetAuthByUsername(username)
	if err != nil || user == nil {
		return ErrInvalidCredentials
	}
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return ErrAccountLocked
	}
	if user.Status == 0 {
		return ErrUserDisabled
	}
	if user.Password != utils.EncodeUserPassword(payload.OldPassword) {
		return ErrInvalidCredentials
	}
	if err := security.ValidatePasswordComplexity(payload.NewPassword); err != nil {
		return err
	}
	return s.userService.ChangeUserPassword(user.ID, payload.NewPassword)
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

func (s *Service) handleLoginFailure(user *model.Auth, ip, reason string) {
	cfg := setting.SecuritySetting
	count := user.FailedLoginCount + 1
	updates := map[string]any{
		"failed_login_count": count,
	}
	if cfg.LoginMaxAttempts > 0 && count >= cfg.LoginMaxAttempts {
		lockDuration := time.Duration(cfg.LoginLockoutMinutes) * time.Minute
		lockUntil := time.Now().Add(lockDuration)
		updates["locked_until"] = lockUntil
		updates["failed_login_count"] = 0
		reason = fmt.Sprintf("%s; locked %d minutes", reason, cfg.LoginLockoutMinutes)
	}
	if _, err := model.Update(&model.Auth{}, map[string]any{"id =": user.ID}, updates); err != nil {
		reason += fmt.Sprintf(" (update error: %v)", err)
	}
	s.logLogin(user.Username, user.ID, ip, false, reason)
}

func (s *Service) logLogin(username string, userID uint, ip string, success bool, message string) {
	status := http.StatusUnauthorized
	if success {
		status = http.StatusOK
	}
	entry := model.AuditLog{
		Category: "login",
		UserID:   userID,
		Username: username,
		IP:       ip,
		Method:   http.MethodPost,
		Path:     "/v1/api/login",
		Status:   status,
		Action:   "login",
		Message:  message,
	}
	go model.CreateAuditLog(entry)
}
