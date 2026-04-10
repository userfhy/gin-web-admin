package security

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-web-admin/utils/setting"
)

func TestValidatePasswordComplexity(t *testing.T) {
	setting.SecuritySetting = &setting.Security{
		PasswordMinLength: 8,
		RequireUppercase:  true,
		RequireLowercase:  true,
		RequireNumber:     true,
		RequireSpecial:    true,
	}

	assert.NoError(t, ValidatePasswordComplexity("GoLang!1"))

	cases := []struct {
		name     string
		password string
	}{
		{"too short", "Aa1!"},
		{"missing upper", "golang!1"},
		{"missing lower", "GOLANG!1"},
		{"missing number", "GoLang!!"},
		{"missing special", "GoLang11"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Error(t, ValidatePasswordComplexity(tt.password))
		})
	}
}
