package security

import (
	"fmt"
	"strings"
	"unicode"

	"gin-web-admin/utils/setting"
)

var specialRunes = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '?', '~', '|', '\\', '/', '.', ',', ':', ';'}

// ValidatePasswordComplexity checks password strength based on security settings.
func ValidatePasswordComplexity(password string) error {
	cfg := setting.SecuritySetting
	if len(password) < cfg.PasswordMinLength {
		return fmt.Errorf("密码长度至少需要 %d 位", cfg.PasswordMinLength)
	}
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasNumber = true
		case unicode.IsSymbol(r), unicode.IsPunct(r), containsRune(specialRunes, r):
			hasSpecial = true
		}
	}
	if cfg.RequireUppercase && !hasUpper {
		return fmt.Errorf("密码需包含至少一个大写字母")
	}
	if cfg.RequireLowercase && !hasLower {
		return fmt.Errorf("密码需包含至少一个小写字母")
	}
	if cfg.RequireNumber && !hasNumber {
		return fmt.Errorf("密码需包含至少一个数字")
	}
	if cfg.RequireSpecial && !hasSpecial {
		return fmt.Errorf("密码需包含至少一个特殊字符(%s)", specialCharset())
	}
	return nil
}

func containsRune(list []rune, target rune) bool {
	for _, r := range list {
		if r == target {
			return true
		}
	}
	return false
}

func specialCharset() string {
	var builder strings.Builder
	for i, r := range specialRunes {
		builder.WriteRune(r)
		if i != len(specialRunes)-1 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}
