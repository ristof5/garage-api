package models

type Service struct {
	ID          int    `json:"id"`
	VehicleID   int    `json:"vehicle_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cost        int    `json:"cost" binding:"required"`
	ServiceDate string `json:"service_date" binding:"required"`
}