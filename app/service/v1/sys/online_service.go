package sysService

import (
	"fmt"
	"strings"
	"time"

	userService "gin-web-admin/app/service/v1/user"
	"gin-web-admin/utils"
	"gin-web-admin/utils/gredis"
)

type OnlineUserQuery struct {
	Pagination utils.Pagination
	Username   string
	IP         string
}

type ForceOfflineStruct struct {
	UserID    uint   `json:"userId" binding:"required"`
	SessionID string `json:"sessionId"`
}

func (s *Service) GetOnlineUserList(query OnlineUserQuery) (utils.PageResult, error) {
	keys, err := gredis.KeysByPrefix("sys:online:user:")
	if err != nil {
		return utils.PageResult{}, err
	}

	userSvc := userService.NewService(s.store)
	list := make([]userService.OnlineSession, 0, len(keys))
	for _, key := range keys {
		item, err := userSvc.GetActiveOnlineSessionByKey(key)
		if err != nil || item == nil {
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
		list = append(list, *item)
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

func (s *Service) ForceOffline(userID uint, sessionID string) error {
	if userID == 0 {
		return fmt.Errorf("user id is required")
	}
	userSvc := userService.NewService(s.store)
	if strings.TrimSpace(sessionID) != "" {
		return userSvc.RemoveOnlineSessionBySessionID(userID, sessionID, true)
	}
	keys, err := gredis.KeysByPrefix(fmt.Sprintf("sys:online:user:%d:", userID))
	if err != nil {
		return err
	}
	activeSessions := make([]userService.OnlineSession, 0, len(keys))
	for _, key := range keys {
		item, getErr := userSvc.GetActiveOnlineSessionByKey(key)
		if getErr != nil {
			return getErr
		}
		if item != nil {
			activeSessions = append(activeSessions, *item)
		}
	}
	if len(activeSessions) > 1 {
		return fmt.Errorf("multiple active sessions found; sessionId is required")
	}
	if len(activeSessions) == 1 && activeSessions[0].SessionID != "" {
		return userSvc.RemoveOnlineSessionBySessionID(userID, activeSessions[0].SessionID, true)
	}
	userSvc.InvalidateAllUserSessions(userID, true)
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
