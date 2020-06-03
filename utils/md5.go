package utils

import (
    "crypto/md5"
    "encoding/hex"
    "gin-test/utils/setting"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
    m := md5.New()
    m.Write([]byte(value))

    return hex.EncodeToString(m.Sum(nil))
}

func EncodeUserPassword(value string) string{
    return EncodeMD5(setting.AppSetting.PasswordSalt + value)
}