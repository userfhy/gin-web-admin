package userService

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/security"
	"strings"
	"time"
)

const (
	onlineSessionKeyPrefix  = "sys:online:user:"
	defaultOnlineSessionTTL = 2 * time.Hour
)

type OnlineSession struct {
	UserID       uint   `json:"userId"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	RoleID       uint   `json:"roleId"`
	RoleName     string `json:"roleName"`
	RoleKey      string `json:"roleKey"`
	IsAdmin      bool   `json:"isAdmin"`
	IP           string `json:"ip"`
	UserAgent    string `json:"userAgent"`
	Browser      string `json:"browser"`
	OS           string `json:"os"`
	LoginAt      string `json:"loginAt"`
	LastActiveAt string `json:"lastActiveAt"`
	ExpiresAt    string `json:"expiresAt"`
	TokenHash    string `json:"tokenHash"`
	Token        string `json:"-"`
}

func onlineSessionKey(userID uint) string {
	return fmt.Sprintf("%s%d", onlineSessionKeyPrefix, userID)
}

func onlineSessionTTL(claims *utils.Claims) time.Duration {
	if claims != nil && claims.ExpiresAt != nil {
		ttl := time.Until(claims.ExpiresAt.Time)
		if ttl > 0 {
			return ttl
		}
	}
	return defaultOnlineSessionTTL
}

func tokenHash(token string) string {
	sum := sha1.Sum([]byte(token))
	return hex.EncodeToString(sum[:])[:16]
}

func (s *Service) HasActiveOnlineSession(userID uint, token string) (bool, error) {
	if userID == 0 || token == "" {
		return false, nil
	}

	session, err := s.GetOnlineSession(userID)
	if err != nil {
		return false, err
	}
	if session == nil {
		return false, nil
	}

	if session.TokenHash == "" {
		return false, nil
	}

	return session.TokenHash == tokenHash(token), nil
}

func parseUserAgent(userAgent string) (browser string, os string) {
	ua := strings.ToLower(strings.TrimSpace(userAgent))
	switch {
	case strings.Contains(ua, "edg/"):
		browser = "Edge"
	case strings.Contains(ua, "chrome/"):
		browser = "Chrome"
	case strings.Contains(ua, "firefox/"):
		browser = "Firefox"
	case strings.Contains(ua, "safari/") && !strings.Contains(ua, "chrome/"):
		browser = "Safari"
	case strings.Contains(ua, "micromessenger"):
		browser = "WeChat"
	default:
		browser = "Unknown"
	}

	switch {
	case strings.Contains(ua, "windows"):
		os = "Windows"
	case strings.Contains(ua, "mac os"):
		os = "macOS"
	case strings.Contains(ua, "android"):
		os = "Android"
	case strings.Contains(ua, "iphone"), strings.Contains(ua, "ipad"), strings.Contains(ua, "ios"):
		os = "iOS"
	case strings.Contains(ua, "linux"):
		os = "Linux"
	default:
		os = "Unknown"
	}
	return
}

func (s *Service) SaveOnlineSession(token string, claims *utils.Claims, ip, userAgent string) error {
	if claims == nil || claims.UserId == 0 || token == "" || gredis.RedisConn == nil {
		return nil
	}

	existing, _ := s.GetOnlineSession(claims.UserId)
	if existing != nil && existing.Token != "" && existing.Token != token {
		s.blockTokenOnly(existing.Token)
	}

	browser, os := parseUserAgent(userAgent)
	now := time.Now()
	loginAt := now.Format(time.RFC3339)
	if existing != nil && existing.LoginAt != "" {
		loginAt = existing.LoginAt
	}

	roleID := uint(0)
	roleName := ""
	nickname := ""
	if existing != nil {
		roleID = existing.RoleID
		roleName = existing.RoleName
		nickname = existing.Nickname
	}
	if existing == nil || existing.Token != token {
		user, err := model.GetUser(map[string]any{"id": claims.UserId})
		if err != nil || user == nil {
			return err
		}
		roleID = user.RoleId
		roleName = user.Role.RoleName
		nickname = user.Nickname
	}

	session := OnlineSession{
		UserID:       claims.UserId,
		Username:     claims.Username,
		Nickname:     nickname,
		RoleID:       roleID,
		RoleName:     roleName,
		RoleKey:      claims.RoleKey,
		IsAdmin:      claims.IsAdmin,
		IP:           security.SanitizePlainText(ip, 64),
		UserAgent:    security.SanitizePlainText(userAgent, 300),
		Browser:      browser,
		OS:           os,
		LoginAt:      loginAt,
		LastActiveAt: now.Format(time.RFC3339),
		ExpiresAt:    expiresAtString(claims, now),
		TokenHash:    tokenHash(token),
		Token:        token,
	}
	return gredis.SetJSON(onlineSessionKey(claims.UserId), session, onlineSessionTTL(claims))
}

func (s *Service) GetOnlineSession(userID uint) (*OnlineSession, error) {
	if userID == 0 || gredis.RedisConn == nil {
		return nil, nil
	}
	var session OnlineSession
	found, err := gredis.GetJSON(onlineSessionKey(userID), &session)
	if err != nil || !found {
		return nil, err
	}
	return &session, nil
}

func (s *Service) RemoveOnlineSession(userID uint) {
	if userID == 0 || gredis.RedisConn == nil {
		return
	}
	if _, err := gredis.Delete(onlineSessionKey(userID)); err != nil {
		logging.Warnf("delete online session failed: %v", err)
	}
}

func (s *Service) blockTokenOnly(jwt string) {
	if jwt == "" {
		return
	}
	if err := s.cacheJWTBlacklist(0, jwt); err != nil {
		logging.Warnf("write online-session blacklist failed: %v", err)
	}
}

func expiresAtString(claims *utils.Claims, now time.Time) string {
	if claims != nil && claims.ExpiresAt != nil {
		return claims.ExpiresAt.Time.Format(time.RFC3339)
	}
	return now.Add(defaultOnlineSessionTTL).Format(time.RFC3339)
}
