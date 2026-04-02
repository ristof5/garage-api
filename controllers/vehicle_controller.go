package controllers

import (
	"net/http"
	"strconv"

	"garage-api/helpers"
	"garage-api/models"
	"garage-api/services"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Service *services.VehicleService
}
// get method, GET ALL VEHICLES
func (vc *VehicleController) GetVehicles(c *gin.Context) {

	vehicles, err := vc.Service.GetAllVehicles()

	if err != nil {

		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch vehicles", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicles fetched successfully", vehicles)
}

//get method, get by id vehicle
func (vc *VehicleController) GetVehicleById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid ID" , err.Error())
		return
	}
	vehicle, err := vc.Service.GetVehicleById(id)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch vehicle", err.Error())
		return
	}
	helpers.SuccessResponse(c, "Vehicle fetched successfully", vehicle)
}
// post method, CREATE A VEHICLE
func (vc *VehicleController) CreateVehicle(c *gin.Context) {

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid request", err.Error())
		return
	}

	err = vc.Service.CreateVehicle(vehicle)

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to create vehicle", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicle created", vehicle)
}

// put method, UPDATE A VEHICLE
func (vc *VehicleController) UpdateVehicle(c *gin.Context) {

	id := c.Param("id")

	var vehicle models.Vehicle

	err := c.ShouldBindJSON(&vehicle)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid request", err.Error())
		return
	}

	err = vc.Service.UpdateVehicle(id, vehicle)

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to update", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Vehicle updated", nil)
}

// delete method, DELETE A VEHICLE
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	id := c.Param("id")

	err := vc.Service.DeleteVehicle(id)

	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete vehicle", err.Error())
		return
	}
	helpers.SuccessResponse(c, "Vehicle deleted successfully", nil)

}

// get method, get vehicle detail by services
func (vc *VehicleController) GetVehicleDetail(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid ID", err.Error())
		return
	}

	vehicle, err := vc.Service.GetVehicleWithServices(id)

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to fetch vehicle detail", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Success get vehicle detail", vehicle)
}