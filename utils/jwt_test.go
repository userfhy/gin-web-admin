package utils

import (
	"testing"
	"time"

	"gin-web-admin/utils/setting"
)

func TestGenerateAndParseToken(t *testing.T) {
	setting.AppSetting = &setting.App{JwtSecret: "unit-test-secret"}
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
	setting.AppSetting = &setting.App{JwtSecret: "unit-test-secret"}
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
