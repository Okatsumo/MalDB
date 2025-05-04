package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()

		log.Printf("[GIN] %d | %v | %s | %-7s %s | UA: %s",
			status,
			duration,
			clientIP,
			method,
			path,
			userAgent,
		)
	}
}
