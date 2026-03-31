package helpers

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"errors":  err,
	})
}