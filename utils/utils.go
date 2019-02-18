package utils

import (
    "log"
    "runtime"
    "encoding/base64"

    "github.com/gin-gonic/gin"
)

func Base64Decode(raw string) ([] byte, error) {
    ecodeBytes, err := base64.StdEncoding.DecodeString(raw)
    if err != nil {
        return nil, err
    }
    return ecodeBytes, nil
}

func PrintStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    log.Printf("==> %s\n", string(buf[:n]))
}

func HandleError(c *gin.Context, err error, http_code int) bool {
    if err != nil {
        _, file, line, ok := runtime.Caller(1)

        log.Printf("Error: %s\nFile: %s Line: %d  %t", err, file, line, ok)

        PrintStack()
        c.JSON(http_code, gin.H{"error": err.Error()})
        return true
    }
    return false
}