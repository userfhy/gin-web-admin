package sysService

import (
	"fmt"
	model "gin-web-admin/app/models"
	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/utils"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
	"strings"
	"time"
)

const (
	jwtBlacklistKeyPrefix = "jwt:blacklist:"
	defaultBlacklistTTL   = 24 * time.Hour
)

type OnlineUserQuery struct {
	Pagination utils.Pagination
	Username   string
	IP         string
}

type ForceOfflineStruct struct {
	UserID uint `json:"userId" binding:"required"`
}

func (s *Service) GetOnlineUserList(query OnlineUserQuery) (utils.PageResult, error) {
	keys, err := gredis.KeysByPrefix("sys:online:user:")
	if err != nil {
		return utils.PageResult{}, err
	}

	list := make([]userService.OnlineSession, 0, len(keys))
	for _, key := range keys {
		var item userService.OnlineSession
		found, err := gredis.GetJSON(key, &item)
		if err != nil || !found {
			continue
		}
		if username := strings.TrimSpace(strings.ToLower(query.Username)); username != "" &&
			!strings.Contains(strings.ToLower(item.Username), username) &&
			!strings.Contains(strings.ToLower(item.Nickname), username) {
			continue
		}
		if ip := strings.TrimSpace(query.IP); ip != "" && !strings.Contains(item.IP, ip) {
			continue
		}
		list = append(list, item)
	}

	sortOnlineSessions(list)

	total := int64(len(list))
	page := query.Pagination.Clone()
	if total == 0 {
		return page.Result([]userService.OnlineSession{}, 0), nil
	}

	offset := page.Offset()
	if offset >= len(list) {
		return page.Result([]userService.OnlineSession{}, total), nil
	}
	end := offset + page.Limit()
	if end > len(list) {
		end = len(list)
	}

	return page.Result(list[offset:end], total), nil
}

func (s *Service) ForceOffline(userID uint) error {
	if userID == 0 {
		return fmt.Errorf("user id is required")
	}

	var session userService.OnlineSession
	found, err := gredis.GetJSON(fmt.Sprintf("sys:online:user:%d", userID), &session)
	if err != nil {
		return err
	}

	if found && session.Token != "" {
		if err := gredis.SetWithTTL(jwtBlacklistKey(session.Token), fmt.Sprintf("%d", userID), blacklistTTL(session.Token)); err != nil {
			logging.Warnf("force offline write blacklist failed: %v", err)
		}
		_ = model.CreateBlockList(userID, session.Token)
	}
	_, _ = model.Update(model.Auth{}, map[string]any{"id =": userID}, map[string]any{"refresh_token": ""})
	userService.InvalidateAuthProfileCache(userID)
	_, _ = gredis.Delete(fmt.Sprintf("sys:online:user:%d", userID))
	return nil
}

func sortOnlineSessions(list []userService.OnlineSession) {
	if len(list) < 2 {
		return
	}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			ti := parseTime(list[i].LastActiveAt)
			tj := parseTime(list[j].LastActiveAt)
			if tj.After(ti) {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func parseTime(value string) time.Time {
	t, _ := time.Parse(time.RFC3339, value)
	return t
}

func jwtBlacklistKey(jwt string) string {
	return jwtBlacklistKeyPrefix + jwt
}

func blacklistTTL(jwt string) time.Duration {
	claims, err := utils.ParseToken(jwt)
	if err != nil || claims == nil || claims.ExpiresAt == nil {
		return defaultBlacklistTTL
	}
	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return time.Minute
	}
	return ttl
}
