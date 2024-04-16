package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"lxtend.com/friend_link/config"
)

var host = fmt.Sprintf("https://%s", config.GetConfig().Domain.Host)

// CorsMiddleware 是处理跨域请求的中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", host)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
