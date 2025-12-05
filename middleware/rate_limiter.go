package middleware

import (
	"backend-kurikulum-apps/dto"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter stores rate limiting data
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int           // requests per window
	window   time.Duration // time window
}

type visitor struct {
	count     int
	lastSeen  time.Time
	resetTime time.Time
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		window:   window,
	}

	// Cleanup goroutine
	go rl.cleanup()

	return rl
}

// cleanup removes old visitors
func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.window)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > rl.window*2 {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// isAllowed checks if request is allowed
func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	v, exists := rl.visitors[ip]

	if !exists {
		rl.visitors[ip] = &visitor{
			count:     1,
			lastSeen:  now,
			resetTime: now.Add(rl.window),
		}
		return true
	}

	// Reset if window has passed
	if now.After(v.resetTime) {
		v.count = 1
		v.resetTime = now.Add(rl.window)
		v.lastSeen = now
		return true
	}

	// Check rate
	if v.count >= rl.rate {
		v.lastSeen = now
		return false
	}

	v.count++
	v.lastSeen = now
	return true
}

// RateLimitMiddleware limits requests per IP
func RateLimitMiddleware(rate int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(rate, window)

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		if !limiter.isAllowed(ip) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, dto.ErrorResponse{
				Success: false,
				Message: "Terlalu banyak request. Silakan coba lagi nanti.",
			})
			return
		}

		ctx.Next()
	}
}
