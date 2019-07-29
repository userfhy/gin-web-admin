package indexController

import (
    "encoding/json"
    "strings"
    
    "gin-test/common"
    "gin-test/utils"
    "github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
    base64 := c.DefaultQuery("base64", "")

    // 替换字符串
    base64String := strings.Replace(base64, " ", "+", -1)

    // base64 解码
    arrByte, error := utils.Base64Decode(base64String)
    if utils.HandleError(c, error, 500) {
        return
    }

    // 结构体
    var imgs []common.ImgText
    err := json.Unmarshal(arrByte, &imgs)

    if utils.HandleError(c, err, 500) {
        return
    }

    c.JSON(200, imgs)
}
