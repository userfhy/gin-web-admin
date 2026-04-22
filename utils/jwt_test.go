package utils

import (
	"testing"
	"time"

	"gin-web-admin/utils/setting"
)

func TestGenerateAndParseToken(t *testing.T) {
	setting.AppSetting = &setting.App{
		JwtSecret:       "unit-test-secret",
		AccessTokenTTL:  setting.Duration(2 * time.Hour),
		RefreshTokenTTL: setting.Duration(7 * 24 * time.Hour),
	}
	claims := Claims{UserId: 9, Username: "tester", RoleKey: "admin", IsAdmin: true}

	start := time.Now()
	token, expire, err := GenerateToken(claims)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	if token == "" {
		t.Fatal("token should not be empty")
	}
	diff := expire.Sub(start)
	if diff < 2*time.Hour-2*time.Second || diff > 2*time.Hour+2*time.Second {
		t.Fatalf("unexpected expire duration %v", diff)
	}

	parsed, err := ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken error: %v", err)
	}
	if parsed.UserId != claims.UserId || parsed.Username != claims.Username || parsed.RoleKey != claims.RoleKey || parsed.IsAdmin != claims.IsAdmin {
		t.Fatalf("parsed claims mismatch: %#v", parsed)
	}

	if _, err := ParseToken(token + "tamper"); err == nil {
		t.Fatal("expected parse error for tampered token")
	}
}

func TestValidateToken(t *testing.T) {
	setting.AppSetting = &setting.App{
		JwtSecret:       "unit-test-secret",
		AccessTokenTTL:  setting.Duration(2 * time.Hour),
		RefreshTokenTTL: setting.Duration(7 * 24 * time.Hour),
	}
	token, _, err := GenerateToken(Claims{UserId: 1, Username: "user", RoleKey: "role", IsAdmin: false})
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken error: %v", err)
	}
	if claims["iss"] != "gin-web-admin" {
		t.Fatalf("issuer mismatch: %v", claims["iss"])
	}

	if _, err := ValidateToken("invalid.token.value"); err == nil {
		t.Fatal("expected validation error for invalid token")
	}
}

func TestGenerateTokenUsesConfiguredTTL(t *testing.T) {
	setting.AppSetting = &setting.App{
		JwtSecret:       "unit-test-secret",
		AccessTokenTTL:  setting.Duration(30 * time.Minute),
		RefreshTokenTTL: setting.Duration(10 * 24 * time.Hour),
	}

	start := time.Now()
	_, expire, err := GenerateToken(Claims{UserId: 1, Username: "user", RoleKey: "role"})
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	diff := expire.Sub(start)
	if diff < 30*time.Minute-2*time.Second || diff > 30*time.Minute+2*time.Second {
		t.Fatalf("unexpected access token ttl %v", diff)
	}

	start = time.Now()
	_, refreshExpire, err := GenerateRefreshToken(Claims{UserId: 1, Username: "user", RoleKey: "role"})
	if err != nil {
		t.Fatalf("GenerateRefreshToken error: %v", err)
	}
	diff = refreshExpire.Sub(start)
	if diff < 10*24*time.Hour-2*time.Second || diff > 10*24*time.Hour+2*time.Second {
		t.Fatalf("unexpected refresh token ttl %v", diff)
	}
}
