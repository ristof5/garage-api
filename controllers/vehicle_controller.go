package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"garage-api/repositories"
)

type VehicleController struct {
	Repo *repositories.VehicleRepository
}

func (vc *VehicleController) GetVehicles(c *gin.Context) {

	vehicles, err := vc.Repo.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, vehicles)
}