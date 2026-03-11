package controllers

import (
	"net/http"

	"garage-api/models"
	"garage-api/repositories"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Repo *repositories.VehicleRepository
}
// get method, GET ALL VEHICLES
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

// post method, CREATE A VEHICLE
func (vc *VehicleController) CreateVehicle(c *gin.Context) {

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	createdVehicle, err := vc.Repo.Create(vehicle)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Vehicle created successfully",
		"vehicle": createdVehicle,
	})
}