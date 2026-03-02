package middlewares

import (
	"Drawer/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(startTime)
		utils.Logger.Info("请求日志",
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"cost", cost.String(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
