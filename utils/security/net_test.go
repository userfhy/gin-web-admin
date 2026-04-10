package security

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-web-admin/utils/setting"
)

func TestIsIPWhitelisted(t *testing.T) {
	setting.SecuritySetting = &setting.Security{
		IPWhitelist: []string{"127.0.0.1/32", "192.168.1.0/24", "10.0.0.5"},
	}
	assert.True(t, IsIPWhitelisted("127.0.0.1"))
	assert.True(t, IsIPWhitelisted("192.168.1.25"))
	assert.True(t, IsIPWhitelisted("10.0.0.5"))
	assert.False(t, IsIPWhitelisted("8.8.8.8"))

	setting.SecuritySetting.IPWhitelist = []string{}
	assert.True(t, IsIPWhitelisted("8.8.8.8"))
}
