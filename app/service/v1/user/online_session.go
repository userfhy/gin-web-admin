package userService

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	model "gin-web-admin/app/models"
	"gin-web-admin/utils"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/security"
)

const (
	onlineSessionKeyPrefix   = "sys:online:user:"
	onlineAccessKeyPrefix    = "sys:online:access:"
	onlineRefreshKeyPrefix   = "sys:online:refresh:"
	defaultOnlineSessionTTL  = 2 * time.Hour
	defaultRefreshSessionTTL = 7 * 24 * time.Hour
)

type OnlineSession struct {
	SessionID             string `json:"sessionId"`
	UserID                uint   `json:"userId"`
	Username              string `json:"username"`
	Nickname              string `json:"nickname"`
	RoleID                uint   `json:"roleId"`
	RoleName              string `json:"roleName"`
	RoleKey               string `json:"roleKey"`
	IsAdmin               bool   `json:"isAdmin"`
	IP                    string `json:"ip"`
	UserAgent             string `json:"userAgent"`
	Browser               string `json:"browser"`
	OS                    string `json:"os"`
	DeviceName            string `json:"deviceName"`
	LoginAt               string `json:"loginAt"`
	LastActiveAt          string `json:"lastActiveAt"`
	ExpiresAt             string `json:"expiresAt"`
	TokenHash             string `json:"tokenHash"`
	TokenRemainingSeconds int64  `json:"tokenRemainingSeconds"`
}

type onlineSessionRecord struct {
	OnlineSession
	RefreshToken     string `json:"refreshToken"`
	RefreshTokenHash string `json:"refreshTokenHash"`
	Token            string `json:"token"`
}

func onlineSessionKey(userID uint, sessionID string) string {
	return fmt.Sprintf("%s%d:%s", onlineSessionKeyPrefix, userID, sessionID)
}

func onlineSessionKeyPrefixByUser(userID uint) string {
	return fmt.Sprintf("%s%d:", onlineSessionKeyPrefix, userID)
}

func onlineAccessKey(token string) string {
	return onlineAccessKeyPrefix + tokenHash(token)
}

func onlineRefreshKey(token string) string {
	return onlineRefreshKeyPrefix + tokenHash(token)
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

func refreshTokenTTL(refreshToken string) time.Duration {
	claims, err := utils.ParseToken(refreshToken)
	if err != nil || claims == nil || claims.ExpiresAt == nil {
		return defaultRefreshSessionTTL
	}
	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return time.Minute
	}
	return ttl
}

func tokenHash(token string) string {
	sum := sha1.Sum([]byte(token))
	return hex.EncodeToString(sum[:])[:16]
}

func sessionIDFromRefresh(refreshToken string) string {
	return tokenHash(refreshToken)
}

func (s *Service) HasActiveOnlineSession(userID uint, token string) (bool, error) {
	if userID == 0 || token == "" {
		return false, nil
	}
	if gredis.RedisConn == nil {
		return true, nil
	}
	session, err := s.GetOnlineSessionByAccessToken(token)
	if err != nil || session == nil {
		return false, err
	}
	return session.UserID == userID && session.TokenHash == tokenHash(token), nil
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

func resolveDeviceName(browser, os string) string {
	browser = strings.TrimSpace(browser)
	os = strings.TrimSpace(os)
	switch {
	case browser != "" && browser != "Unknown" && os != "" && os != "Unknown":
		return os + " / " + browser
	case os != "" && os != "Unknown":
		return os
	case browser != "" && browser != "Unknown":
		return browser
	default:
		return "Unknown"
	}
}

func (s *Service) CreateLoginSession(accessToken, refreshToken string, claims *utils.Claims, ip, userAgent string) error {
	if claims == nil || claims.UserId == 0 || accessToken == "" || refreshToken == "" || gredis.RedisConn == nil {
		return nil
	}

	profile, err := s.GetAuthProfile(claims.UserId)
	if err != nil || profile == nil {
		return err
	}

	now := time.Now()
	browser, os := parseUserAgent(userAgent)
	session := onlineSessionRecord{
		OnlineSession: OnlineSession{
			SessionID:    sessionIDFromRefresh(refreshToken),
			UserID:       claims.UserId,
			Username:     claims.Username,
			Nickname:     profile.Nickname,
			RoleID:       profile.RoleID,
			RoleName:     profile.RoleName,
			RoleKey:      claims.RoleKey,
			IsAdmin:      claims.IsAdmin,
			IP:           security.SanitizePlainText(ip, 64),
			UserAgent:    security.SanitizePlainText(userAgent, 300),
			Browser:      browser,
			OS:           os,
			DeviceName:   resolveDeviceName(browser, os),
			LoginAt:      now.Format(time.RFC3339),
			LastActiveAt: now.Format(time.RFC3339),
			ExpiresAt:    expiresAtString(claims, now),
			TokenHash:    tokenHash(accessToken),
		},
		RefreshTokenHash: tokenHash(refreshToken),
		Token:            accessToken,
		RefreshToken:     refreshToken,
	}
	return s.persistOnlineSession(nil, session, claims)
}

func (s *Service) RefreshOnlineSessionAccessToken(accessToken, refreshToken string, claims *utils.Claims, ip, userAgent string) error {
	if claims == nil || claims.UserId == 0 || accessToken == "" || refreshToken == "" || gredis.RedisConn == nil {
		return nil
	}

	existing, err := s.getOnlineSessionRecordByRefreshToken(refreshToken)
	if err != nil {
		return err
	}
	if existing == nil || existing.UserID != claims.UserId {
		return fmt.Errorf("该access_token对应的用户信息不存在")
	}

	now := time.Now()
	browser, os := parseUserAgent(userAgent)
	existing.Token = accessToken
	existing.TokenHash = tokenHash(accessToken)
	existing.ExpiresAt = expiresAtString(claims, now)
	existing.LastActiveAt = now.Format(time.RFC3339)
	if sanitizedIP := security.SanitizePlainText(ip, 64); sanitizedIP != "" {
		existing.IP = sanitizedIP
	}
	if sanitizedUA := security.SanitizePlainText(userAgent, 300); sanitizedUA != "" {
		existing.UserAgent = sanitizedUA
		existing.Browser = browser
		existing.OS = os
		existing.DeviceName = resolveDeviceName(browser, os)
	}
	return s.persistOnlineSession(existing, *existing, claims)
}

func (s *Service) TouchOnlineSession(accessToken string, claims *utils.Claims, ip, userAgent string) error {
	if claims == nil || claims.UserId == 0 || accessToken == "" || gredis.RedisConn == nil {
		return nil
	}

	session, err := s.getOnlineSessionRecordByAccessToken(accessToken)
	if err != nil || session == nil {
		return err
	}

	now := time.Now()
	session.LastActiveAt = now.Format(time.RFC3339)
	session.ExpiresAt = expiresAtString(claims, now)
	if sanitizedIP := security.SanitizePlainText(ip, 64); sanitizedIP != "" {
		session.IP = sanitizedIP
	}
	if sanitizedUA := security.SanitizePlainText(userAgent, 300); sanitizedUA != "" {
		browser, os := parseUserAgent(userAgent)
		session.UserAgent = sanitizedUA
		session.Browser = browser
		session.OS = os
		session.DeviceName = resolveDeviceName(browser, os)
	}
	session.Token = accessToken
	session.TokenHash = tokenHash(accessToken)
	return s.persistOnlineSession(session, *session, claims)
}

func (s *Service) persistOnlineSession(previous *onlineSessionRecord, session onlineSessionRecord, claims *utils.Claims) error {
	refreshTTL := refreshTokenTTL(session.RefreshToken)
	accessTTL := onlineSessionTTL(claims)
	sessionKey := onlineSessionKey(session.UserID, session.SessionID)

	if previous != nil && previous.Token != "" && previous.Token != session.Token {
		_, _ = gredis.Delete(onlineAccessKey(previous.Token))
	}

	if err := gredis.SetJSON(sessionKey, session, refreshTTL); err != nil {
		return err
	}
	if err := gredis.SetString(onlineRefreshKey(session.RefreshToken), sessionKey, refreshTTL); err != nil {
		return err
	}
	if err := gredis.SetString(onlineAccessKey(session.Token), sessionKey, accessTTL); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetOnlineSessionByAccessToken(accessToken string) (*OnlineSession, error) {
	record, err := s.getOnlineSessionRecordByAccessToken(accessToken)
	if err != nil || record == nil {
		return nil, err
	}
	return &record.OnlineSession, nil
}

func (s *Service) GetActiveOnlineSessionByKey(sessionKey string) (*OnlineSession, error) {
	record, err := s.getOnlineSessionRecordByKey(sessionKey)
	if err != nil || record == nil {
		return nil, err
	}
	active, err := s.isSessionAccessActive(record)
	if err != nil || !active {
		return nil, err
	}
	fillSessionRuntimeFields(&record.OnlineSession)
	return &record.OnlineSession, nil
}

func (s *Service) getOnlineSessionRecordByAccessToken(accessToken string) (*onlineSessionRecord, error) {
	if accessToken == "" || gredis.RedisConn == nil {
		return nil, nil
	}
	sessionKey, err := gredis.GetString(onlineAccessKey(accessToken))
	if err != nil {
		return nil, nil
	}
	return s.getOnlineSessionRecordByKey(sessionKey)
}

func (s *Service) GetOnlineSessionByRefreshToken(refreshToken string) (*OnlineSession, error) {
	record, err := s.getOnlineSessionRecordByRefreshToken(refreshToken)
	if err != nil || record == nil {
		return nil, err
	}
	return &record.OnlineSession, nil
}

func (s *Service) getOnlineSessionRecordByRefreshToken(refreshToken string) (*onlineSessionRecord, error) {
	if refreshToken == "" || gredis.RedisConn == nil {
		return nil, nil
	}
	sessionKey, err := gredis.GetString(onlineRefreshKey(refreshToken))
	if err != nil {
		return nil, nil
	}
	return s.getOnlineSessionRecordByKey(sessionKey)
}

func (s *Service) getOnlineSessionRecordByKey(sessionKey string) (*onlineSessionRecord, error) {
	if sessionKey == "" || gredis.RedisConn == nil {
		return nil, nil
	}
	var session onlineSessionRecord
	found, err := gredis.GetJSON(sessionKey, &session)
	if err != nil || !found {
		return nil, err
	}
	return &session, nil
}

func (s *Service) RemoveOnlineSessionByAccessToken(userID uint, accessToken string) {
	if accessToken == "" || gredis.RedisConn == nil {
		return
	}
	session, err := s.getOnlineSessionRecordByAccessToken(accessToken)
	if err != nil || session == nil {
		return
	}
	if userID > 0 && session.UserID != userID {
		return
	}
	s.deleteOnlineSession(*session)
}

func (s *Service) RemoveOnlineSessionBySessionID(userID uint, sessionID string, blockAccess bool) error {
	if userID == 0 || sessionID == "" || gredis.RedisConn == nil {
		return nil
	}
	session, err := s.getOnlineSessionRecordByKey(onlineSessionKey(userID, sessionID))
	if err != nil || session == nil {
		return err
	}
	if blockAccess && session.Token != "" {
		s.blockTokenOnly(session.Token)
		_ = model.CreateBlockList(userID, session.Token)
	}
	s.deleteOnlineSession(*session)
	return nil
}

func (s *Service) InvalidateAllUserSessions(userID uint, blockAccess bool) {
	if userID == 0 || gredis.RedisConn == nil {
		return
	}
	keys, err := gredis.KeysByPrefix(onlineSessionKeyPrefixByUser(userID))
	if err != nil {
		logging.Warnf("scan online sessions failed: %v", err)
		return
	}
	for _, key := range keys {
		session, getErr := s.getOnlineSessionRecordByKey(key)
		if getErr != nil || session == nil {
			continue
		}
		if blockAccess && session.Token != "" {
			s.blockTokenOnly(session.Token)
			_ = model.CreateBlockList(userID, session.Token)
		}
		s.deleteOnlineSession(*session)
	}
}

func (s *Service) deleteOnlineSession(session onlineSessionRecord) {
	keys := []string{onlineSessionKey(session.UserID, session.SessionID)}
	if session.Token != "" {
		keys = append(keys, onlineAccessKey(session.Token))
	}
	if session.RefreshToken != "" {
		keys = append(keys, onlineRefreshKey(session.RefreshToken))
	}
	if _, err := gredis.DeleteKeys(keys...); err != nil {
		logging.Warnf("delete online session failed: %v", err)
	}
}

func (s *Service) isSessionAccessActive(session *onlineSessionRecord) (bool, error) {
	if session == nil || session.Token == "" {
		return false, nil
	}
	if isSessionExpired(session.ExpiresAt) {
		return false, nil
	}
	return gredis.ExistsKey(onlineAccessKey(session.Token))
}

func isSessionExpired(expiresAt string) bool {
	if strings.TrimSpace(expiresAt) == "" {
		return true
	}
	expireAtTime, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		return true
	}
	return !expireAtTime.After(time.Now())
}

func fillSessionRuntimeFields(session *OnlineSession) {
	if session == nil {
		return
	}
	if strings.TrimSpace(session.DeviceName) == "" {
		session.DeviceName = resolveDeviceName(session.Browser, session.OS)
	}
	expireAtTime, err := time.Parse(time.RFC3339, session.ExpiresAt)
	if err != nil {
		session.TokenRemainingSeconds = 0
		return
	}
	remaining := time.Until(expireAtTime)
	if remaining <= 0 {
		session.TokenRemainingSeconds = 0
		return
	}
	session.TokenRemainingSeconds = int64(remaining.Seconds())
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
		return claims.ExpiresAt.Format(time.RFC3339)
	}
	return now.Add(defaultOnlineSessionTTL).Format(time.RFC3339)
}
