package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// 错误响应 (code 暂时未有自定义)
func Error(c *gin.Context, httpCode int, code int, msg string) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: msg,
	})
}

// 默认响应 (code 暂定为 1)
func ErrorDefault(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: msg,
	})
}
