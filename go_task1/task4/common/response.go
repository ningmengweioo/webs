package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// {“code”: 200, "data": value}的形式响应
func Success(c *gin.Context, value interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": value,
		"msg":  "success",
	})
}

// ParamsErrorResp 参数错误
func ParamsErrorRes(c *gin.Context, msg string, value ...interface{}) {
	//config.Logger(c).Warn(fmt.Sprintf("参数错误 %v", value))
	//若是msg 为空,就默认值"参数错误"
	if msg == "" {
		msg = "参数错误"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  500,
		"msg":   msg,
		"error": value,
	})
}

// UnknownErrorResp 未知错误、服务器错误
func UnknownErrorRes(c *gin.Context, msg string, err ...error) {
	//config.Logger(c).Error(err, err.Error())
	var errorMsg interface{}

	// 检查是否提供了err参数
	if len(err) > 0 && err[0] != nil {
		errorMsg = err[0].Error()
	} else {
		errorMsg = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  500,
		"msg":   msg,
		"error": errorMsg,
	})

}
