package userService

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/utils/gredis"
	"strconv"
	"time"
)

const (
	authProfileCachePrefix = "sys:auth:profile:"
	authProfileCacheTTL    = 24 * time.Hour
)

type AuthProfileCache struct {
	UserID       uint   `json:"userId"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	RoleID       uint   `json:"roleId"`
	RoleName     string `json:"roleName"`
	RoleKey      string `json:"roleKey"`
	IsAdmin      bool   `json:"isAdmin"`
	Status       int    `json:"status"`
	Password     string `json:"password"`
	RefreshToken string `json:"refreshToken"`
	LastLoginIP  string `json:"lastLoginIP"`
}

func authProfileCacheKey(userID uint) string {
	return authProfileCachePrefix + formatUint(userID)
}

func formatUint(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

func buildAuthProfileCache(user *model.Auth) *AuthProfileCache {
	if user == nil {
		return nil
	}
	return &AuthProfileCache{
		UserID:       user.ID,
		Username:     user.Username,
		Nickname:     user.Nickname,
		RoleID:       user.RoleId,
		RoleName:     user.Role.RoleName,
		RoleKey:      user.Role.RoleKey,
		IsAdmin:      user.Role.IsAdmin,
		Status:       user.Status,
		Password:     user.Password,
		RefreshToken: user.RefreshToken,
		LastLoginIP:  user.LastLoginIP,
	}
}

func (s *Service) CacheAuthProfile(user *model.Auth) {
	if user == nil || user.ID == 0 || gredis.RedisConn == nil {
		return
	}
	profile := buildAuthProfileCache(user)
	if profile == nil {
		return
	}
	gredis.SetJSONAsync(authProfileCacheKey(user.ID), profile, authProfileCacheTTL)
}

func (s *Service) GetAuthProfile(userID uint) (*AuthProfileCache, error) {
	if userID == 0 {
		return nil, nil
	}
	if gredis.RedisConn != nil {
		var cached AuthProfileCache
		found, err := gredis.GetJSON(authProfileCacheKey(userID), &cached)
		if err == nil && found {
			return &cached, nil
		}
		if err != nil {
			return nil, err
		}
	}

	user, err := model.GetUser(map[string]any{"id": userID})
	if err != nil || user == nil {
		return nil, err
	}

	profile := buildAuthProfileCache(user)
	if profile != nil && gredis.RedisConn != nil {
		gredis.SetJSONAsync(authProfileCacheKey(userID), profile, authProfileCacheTTL)
	}
	return profile, nil
}

func (s *Service) InvalidateAuthProfile(userID uint) {
	invalidateAuthProfile(userID)
}

func invalidateAuthProfile(userID uint) {
	if userID == 0 || gredis.RedisConn == nil {
		return
	}
	_, _ = gredis.Delete(authProfileCacheKey(userID))
}

func InvalidateAuthProfileCache(userID uint) {
	invalidateAuthProfile(userID)
}

func (s *Service) UpdateAuthProfileRefreshToken(userID uint, refreshToken string) {
	if userID == 0 || gredis.RedisConn == nil {
		return
	}
	profile, err := s.GetAuthProfile(userID)
	if err != nil || profile == nil {
		return
	}
	profile.RefreshToken = refreshToken
	gredis.SetJSONAsync(authProfileCacheKey(userID), profile, authProfileCacheTTL)
}

func InvalidateAllAuthProfileCache() {
	gredis.DeleteByPrefixAsync(authProfileCachePrefix)
}
