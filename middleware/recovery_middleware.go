package middleware

import (
	"backend-kurikulum-apps/dto"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware recovers from panics
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				log.Printf("Panic recovered: %v\nStack trace:\n%s", err, debug.Stack())

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
					Success: false,
					Message: "Terjadi kesalahan internal server",
				})
			}
		}()

		ctx.Next()
	}
}
