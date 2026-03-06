package main

import (
	"garage-api/config"
	"garage-api/controllers"
	"garage-api/repositories"
	"garage-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.ConnectDB()

	if err != nil {
		panic(err)
	}

	repo := repositories.VehicleRepository{DB: db}
	controller := controllers.VehicleController{Repo: &repo}

	r := gin.Default()

	routes.SetupRoutes(r, &controller)

	r.Run(":8080")
}