package middleware

import (
	"github.com/gin-gonic/gin"
	"lxtend.com/friend_link/config"
	"lxtend.com/friend_link/logger"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.GetConfig().Auth.Key
		if c.GetHeader("Authorization") != secret {
			logger.Warn("Unauthorized request")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
