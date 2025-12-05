package middleware

import (
	"backend-kurikulum-apps/config"
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware(jwtConfig *config.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Authorization header diperlukan",
			})
			return
		}

		// Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Format authorization tidak valid",
			})
			return
		}

		token := parts[1]
		claims, err := helper.ValidateJWT(token, jwtConfig.Secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Token tidak valid atau sudah expired",
				Error:   err.Error(),
			})
			return
		}

		// Set user info to context
		ctx.Set("userID", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Unauthorized",
			})
			return
		}

		userRole := role.(string)
		allowed := false
		for _, r := range allowedRoles {
			if userRole == r {
				allowed = true
				break
			}
		}

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, dto.ErrorResponse{
				Success: false,
				Message: "Anda tidak memiliki akses ke resource ini",
			})
			return
		}

		ctx.Next()
	}
}

// KaprodiOnly middleware - only kaprodi can access
func KaprodiOnly() gin.HandlerFunc {
	return RoleMiddleware("kaprodi")
}

// DosenOnly middleware - only dosen can access
func DosenOnly() gin.HandlerFunc {
	return RoleMiddleware("dosen")
}

// KaprodiOrDosen middleware - kaprodi or dosen can access
func KaprodiOrDosen() gin.HandlerFunc {
	return RoleMiddleware("kaprodi", "dosen")
}

// OptionalAuth middleware - auth is optional
func OptionalAuth(jwtConfig *config.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Next()
			return
		}

		token := parts[1]
		claims, err := helper.ValidateJWT(token, jwtConfig.Secret)
		if err == nil {
			ctx.Set("userID", claims.UserID)
			ctx.Set("email", claims.Email)
			ctx.Set("role", claims.Role)
		}

		ctx.Next()
	}
}
