package utils

import (
	"testing"

	"gin-web-admin/utils/setting"
)

func TestEncodeMD5(t *testing.T) {
	got := EncodeMD5("gin")
	const want = "a6c72983f8a0a002155d67b12b345629"
	if got != want {
		t.Fatalf("EncodeMD5 mismatch got %s want %s", got, want)
	}
}

func TestEncodeUserPassword(t *testing.T) {
	setting.AppSetting = &setting.App{PasswordSalt: "salt-"}

	got := EncodeUserPassword("secret")
	want := EncodeMD5("salt-secret")
	if got != want {
		t.Fatalf("EncodeUserPassword should md5 salt+password, got %s want %s", got, want)
	}
}
