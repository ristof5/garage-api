package models

type VehicleDetail struct {
	ID       int       `json:"id"`
	Brand    string    `json:"brand"`
	Model    string    `json:"model"`
	Year     int       `json:"year"`
	Services []Service `json:"services"`
}