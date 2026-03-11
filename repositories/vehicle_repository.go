package repositories

import (
	"database/sql"
	"garage-api/models"
)

type VehicleRepository struct {
	DB *sql.DB
}

// GET ALL VEHICLES
func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {

	rows, err := r.DB.Query("SELECT id, brand, model, year FROM vehicles")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vehicles []models.Vehicle

	for rows.Next() {

		var v models.Vehicle

		err := rows.Scan(&v.ID, &v.Brand, &v.Model, &v.Year)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}

// CREATE A VEHICLE
func (r *VehicleRepository) Create(vehicle models.Vehicle) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO vehicles (brand, model, year) VALUES (?, ?, ?)", vehicle.Brand, vehicle.Model, vehicle.Year)//values with ? to prevent sql injection
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

