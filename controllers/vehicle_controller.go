package controllers

import (
	"net/http"

	"garage-api/helpers"
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

		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch vehicles", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicles fetched successfully", vehicles)
}

// post method, CREATE A VEHICLE
func (vc *VehicleController) CreateVehicle(c *gin.Context) {

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	_, err = vc.Repo.Create(vehicle)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create vehicle", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicle created successfully", vehicle)
}

// put method, UPDATE A VEHICLE
func (vc *VehicleController) UpdateVehicle(c *gin.Context) {

	id := c.Param("id")

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)

	if err != nil {

		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err = vc.Repo.UpdateVehicle(id, vehicle)

	if err != nil {

		helpers.ErrorResponse(c, http.StatusBadRequest, "Failed to update Vehicle", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicle updated successfully", vehicle)
}

// delete method, DELETE A VEHICLE
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	id := c.Param("id")

	err := vc.Repo.DeleteVehicle(id)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete vehicle", err.Error())
		return
	}
	helpers.SuccessResponse(c, "Vehicle deleted successfully", nil)

}