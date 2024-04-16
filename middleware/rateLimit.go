package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"lxtend.com/friend_link/logger"
	"lxtend.com/friend_link/utils"
)

type rateLimitConfig struct {
	MaxRequests    int
	Window         int
	RequestHistory map[string]int
}

var rateLimit rateLimitConfig

var tokenBucket utils.TokenBucket

func init() {
	rateLimit.MaxRequests = 100
	rateLimit.Window = 60
	rateLimit.RequestHistory = make(map[string]int)
	tokenBucket = *utils.NewTokenBucket(100, 20)
	tokenBucket.StartAddToken(200 * time.Millisecond)
}

func RateLimitMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ip := c.ClientIP()
		rateLimit.RequestHistory[ip]++
		logger.Info("Request from IP: ", ip, " count: ", rateLimit.RequestHistory[ip], " avaliable token: ", tokenBucket.GetTokenCount())
		if !tokenBucket.GetToken() {
			logger.Warn("Too many requests now!!! Current ip:", ip)
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	})
}
