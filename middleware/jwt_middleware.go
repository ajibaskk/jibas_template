// middleware/jwt_middleware.go
package middleware

import (
	"jibas-template/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is the middleware function for protecting routes with JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Validate token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context for use in handlers
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
