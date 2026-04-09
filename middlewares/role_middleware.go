package middlewares

import (
	"net/http"

	"garage-api/helpers"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			helpers.ErrorResponse(c, http.StatusForbidden, "Role not found in token", nil)
			c.Abort()
			return
		}

		role, ok := roleVal.(string)
		if !ok {
			helpers.ErrorResponse(c, http.StatusForbidden, "Invalid role format", nil)
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		helpers.ErrorResponse(c, http.StatusForbidden, "Access denied: insufficient role", nil)
		c.Abort()
	}
}