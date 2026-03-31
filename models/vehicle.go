package models

type Vehicle struct {
	ID    int    `json:"id"`
	Brand string `json:"brand" binding:"required"`
	Model string `json:"model" binding:"required"`
	Year  int    `json:"year" binding:"required"`
}