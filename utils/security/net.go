package security

import (
	"net"
	"strings"

	"gin-web-admin/utils/setting"
)

// IsIPWhitelisted checks if ip belongs to configured whitelist; returns true when whitelist empty.
func IsIPWhitelisted(ip string) bool {
	whitelist := setting.SecuritySetting.IPWhitelist
	if len(whitelist) == 0 {
		return true
	}
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	for _, item := range whitelist {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if strings.Contains(item, "/") {
			if _, network, err := net.ParseCIDR(item); err == nil && network.Contains(parsedIP) {
				return true
			}
			continue
		}
		if parsedIP.Equal(net.ParseIP(item)) {
			return true
		}
	}
	return false
}
