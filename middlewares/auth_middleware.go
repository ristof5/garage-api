package middlewares

import (
	"net/http"
	"os"
	"strings"

	"garage-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		

		if authHeader == "" {
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", nil)
			c.Abort()
			return
		}

		// format: Bearer TOKEN
		tokenString := strings.Split(authHeader, " ")

		if len(tokenString) != 2 {
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid token format", nil)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}