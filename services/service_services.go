package services

import (
	"garage-api/models"
	"garage-api/repositories"
)

type ServiceService struct {
	Repo *repositories.ServiceRepository
}

func (s *ServiceService) CreateService(service models.Service) error {
	return s.Repo.Create(service)
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	return s.Repo.GetAll()
}

func (s *ServiceService) GetServicesByVehicleID(vehicleId int) ([]models.Service, error){
	return s.Repo.GetByVehicleID(vehicleId)
}