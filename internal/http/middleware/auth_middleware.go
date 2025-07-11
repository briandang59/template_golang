package middleware

import (
	"net/http"
	"strings"

	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/http/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, "missing or invalid Authorization header")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecret()), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "invalid token claims")
			c.Abort()
			return
		}

		// Chuyển kiểu an toàn
		idFloat, ok1 := claims["id"].(float64)
		username, ok2 := claims["username"].(string)
		if !ok1 || !ok2 {
			response.Error(c, http.StatusUnauthorized, "invalid token content")
			c.Abort()
			return
		}

		accountID := uint(idFloat)

		// Lưu vào context để dùng trong handler
		c.Set("account_id", accountID)
		c.Set("username", username)

		c.Next()
	}
}
