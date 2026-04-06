package controllers

import (
	// "net/http"

	"garage-api/helpers"
	"garage-api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func (ac *AuthController) Login(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&req)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid request", err.Error())
		return
	}

	token, err := ac.Service.Login(req.Username, req.Password)

	if err != nil {
		helpers.ErrorResponse(c, 401, err.Error(), nil)
		return
	}

	helpers.SuccessResponse(c, "Login success", gin.H{
		"token": token,
	})
}