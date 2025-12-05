package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs request and response
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Get request body
		var requestBody []byte
		if ctx.Request.Body != nil {
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Process request
		ctx.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Log format
		log.Printf("[%s] %s %s | Status: %d | Latency: %v | IP: %s | User-Agent: %s",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Request.URL.RawQuery,
			ctx.Writer.Status(),
			latency,
			ctx.ClientIP(),
			ctx.Request.UserAgent(),
		)
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		ctx.Set("requestID", requestID)
		ctx.Header("X-Request-ID", requestID)
		ctx.Next()
	}
}

func generateRequestID() string {
	return time.Now().Format("20060102150405") + randomString(6)
}

func randomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[time.Now().UnixNano()%int64(len(letterBytes))]
	}
	return string(b)
}
