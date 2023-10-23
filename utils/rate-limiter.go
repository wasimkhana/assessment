package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
	"time"
)

type IPRateLimiter struct {
	requestCount map[string]int
	lastReset    time.Time
	mutex        sync.Mutex
}

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		requestCount: make(map[string]int),
		lastReset:    time.Now().Truncate(24 * time.Hour),
	}
}

func (limiter *IPRateLimiter) LimitRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := strings.Split(c.ClientIP(), ":")[0]

		limiter.mutex.Lock()
		defer limiter.mutex.Unlock()

		now := time.Now()
		if now.Day() != limiter.lastReset.Day() {
			limiter.requestCount = make(map[string]int)
			limiter.lastReset = now.Truncate(24 * time.Hour)
		}

		if limiter.requestCount[clientIP] >= 5 {
			c.JSON(http.StatusTooManyRequests, gin.H{"status": 429, "error": "Rate Limit Exceeded", "message": "Upgrade your account for higher rate limits", "data": nil})
			c.Abort()
			return
		}

		limiter.requestCount[clientIP]++
		c.Next()
	}
}
