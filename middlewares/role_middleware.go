package middlewares

import (
	"net/http"
	"garage-api/helpers"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists{
			helpers.ErrorResponse(c, http.StatusForbidden, "Role not found", nil)
			c.Abort()
			return 
		}
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return 
			}
		}
		helpers.ErrorResponse(c, http.StatusForbidden, "Access Denied", nil)
		c.Abort()
	}
}