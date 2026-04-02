package repositories

import (
	"database/sql"
	"garage-api/models"
)

type ServiceRepository struct {
	DB *sql.DB
}

// POST method to CREATE A SERVICE
func (r *ServiceRepository) Create(service models.Service) error {
	query := `INSERT INTO services (vehicle_id, description, cost, service_date) VALUES (?, ?, ?, ?)`
	_, err := r.DB.Exec(query, service.VehicleID, service.Description, service.Cost, service.ServiceDate)
	return err
}

// GET method to get all Service
func (r *ServiceRepository) GetAll() ([]models.Service, error) {
	rows, err := r.DB.Query(`
		SELECT id, vehicle_id, description, cost, service_date FROM services
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var s models.Service

		err := rows.Scan(
			&s.ID,
			&s.VehicleID,
			&s.Description,
			&s.Cost,
			&s.ServiceDate,
		)

		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

// GET method to get services by id
func (r *ServiceRepository) GetByVehicleID(vehicleID int) ([]models.Service, error) {

	query := `
	SELECT id, vehicle_id, description, cost, service_date 
	FROM services
	WHERE vehicle_id = ?
	`

	rows, err := r.DB.Query(query, vehicleID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var s models.Service

		err := rows.Scan(
			&s.ID,
			&s.VehicleID,
			&s.Description,
			&s.Cost,
			&s.ServiceDate,
		)

		if err != nil {
			return nil, err 
		}

		services = append(services, s)
	}

	return services, nil
}

// put method to update service
// func (r *ServiceRepository) UpdateServices() (models.Service, error) {

// }