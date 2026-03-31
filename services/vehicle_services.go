package services

import (
	"errors"
	"garage-api/models"
	"garage-api/repositories"
)

	type VehicleService struct {
		Repo *repositories.VehicleRepository
	}

	func (s *VehicleService) GetAllVehicles() ([]models.Vehicle, error) {
		return s.Repo.GetAll()
	}

	func (s *VehicleService) CreateVehicle(vehicle models.Vehicle) error {
		if vehicle.Year < 1900 {
			return errors.new("year is not valid")
		}
		return s.Repo.Create(vehicle)
	}

	func (s *VehicleService) UpdateVehicle(id string, vehicle models.Vehicle) error {
		if vehicle.Brand == "" {
			return errors.new("brand cannot be empty")
		}
		return s.Repo.UpdateVehicle(id, vehicle)
	}

	func (s *VehicleService) DeleteVehicle(id string) error{
		return s.Repo.DeleteVehicle(id)
	}