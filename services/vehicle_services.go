package services

import (
	"errors"
	"garage-api/models"
	"garage-api/repositories"
)

	type VehicleService struct {
		Repo *repositories.VehicleRepository
	}
	
 	// get all vehicles
	func (s *VehicleService) GetAllVehicles() ([]models.Vehicle, error) {
		return s.Repo.GetAll()
	}
	// get vehicle by id
	func (s *VehicleService) GetVehicleById(id int) (models.Vehicle, error) {
		if id <= 0 {
			return models.Vehicle{}, errors.New("invalid id")
		}
		vehicle, err := s.Repo.GetVehicleById(id)
		if  err != nil {
			return models.Vehicle{}, err
		}
		return vehicle, nil
	}

	func (s *VehicleService) CreateVehicle(vehicle models.Vehicle) error {
		if vehicle.Year < 1900 {
			return errors.New("year is not valid")
		}
		_, err := s.Repo.Create(vehicle)
		return err
	}

	func (s *VehicleService) UpdateVehicle(id string, vehicle models.Vehicle) error {
		if vehicle.Brand == "" {
			return errors.New("brand cannot be empty")
		}
		return s.Repo.UpdateVehicle(id, vehicle)
	}

	func (s *VehicleService) DeleteVehicle(id string) error{
		return s.Repo.DeleteVehicle(id)
	}

	func (s *VehicleService) GetVehicleWithServices(id int) (models.VehicleDetail, error) {
	return s.Repo.GetVehicleWithServices(id)
}