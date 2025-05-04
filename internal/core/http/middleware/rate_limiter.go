package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

type ClientLimiter struct {
	clients map[string]*rate.Limiter
	mu      sync.Mutex
	r       rate.Limit
	burst   int
}

func NewClientLimiter(r int, b int) *ClientLimiter {
	return &ClientLimiter{
		clients: make(map[string]*rate.Limiter),
		r:       rate.Limit(r),
		burst:   b,
	}
}

func (c *ClientLimiter) getLimiter(ip string) *rate.Limiter {
	c.mu.Lock()
	defer c.mu.Unlock()

	limiter, exists := c.clients[ip]
	if !exists {
		limiter = rate.NewLimiter(c.r, c.burst)
		c.clients[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(limiter *ClientLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		limiter := limiter.getLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}

		ctx.Next()
	}
}
