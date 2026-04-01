package controllers

import (
	// "net/http"
	"strconv"

	"garage-api/helpers"
	"garage-api/models"
	"garage-api/services"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	Service *services.ServiceService
}

func (sc *ServiceController) CreateService(c *gin.Context) {

	var service models.Service

	err := c.ShouldBindJSON(&service)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid request", err.Error())
		return
	}

	err = sc.Service.CreateService(service)

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to create service", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Service created", service)
}

func (sc *ServiceController) GetServices(c *gin.Context) {

	services, err := sc.Service.GetAllServices()

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to fetch services", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Success get services", services)
}

func (sc *ServiceController) GetServicesByVehicle(c *gin.Context) {

	idParam := c.Param("id")

	vehicleID, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(c, 400, "Invalid vehicle ID", err.Error())
		return
	}

	services, err := sc.Service.GetServicesByVehicleID(vehicleID)

	if err != nil {
		helpers.ErrorResponse(c, 500, "Failed to fetch services", err.Error())
		return
	}

	helpers.SuccessResponse(c, "Success get vehicle services", services)
}