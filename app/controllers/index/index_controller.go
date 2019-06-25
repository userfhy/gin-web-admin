package index_controller

import (
    "encoding/json"
    "strings"

    "github.com/fanhengyuan/gin-test/common"
    "github.com/fanhengyuan/gin-test/utils"
    "github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
    base64 := c.DefaultQuery("base64", "")

    // 替换字符串
    base64_string := strings.Replace(base64, " ", "+", -1)

    // base64 解码
    arr_byte, error := utils.Base64Decode(base64_string)
    if utils.HandleError(c, error, 500) {
        return
    }

    // 结构体
    var imgs []common.ImgText
    err := json.Unmarshal(arr_byte, &imgs)

    if utils.HandleError(c, err, 500) {
        return
    }

    c.JSON(200, imgs)
}
