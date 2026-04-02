package repositories

import (
	"database/sql"
	"fmt"
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

// get vehicle by id
func (r *VehicleRepository) GetVehicleById(id int) (models.Vehicle, error) {
	var vehicle models.Vehicle

	query := "SELECT id, brand, model, year FROM vehicles WHERE id = ?"

	row := r.DB.QueryRow(query, id)

	err := row.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return vehicle, fmt.Errorf("vehicle not found")
		}
		return vehicle, err
	}
	return vehicle, nil
}

// CREATE A VEHICLE
func (r *VehicleRepository) Create(vehicle models.Vehicle) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO vehicles (brand, model, year) VALUES (?, ?, ?)", vehicle.Brand, vehicle.Model, vehicle.Year)//values with ? to prevent sql injection
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// UPDATE A VEHICLE
func (r *VehicleRepository) UpdateVehicle(id string, vehicle models.Vehicle) error {

	query := `
	UPDATE vehicles 
	SET brand = ?, model = ?, year = ?
	WHERE id = ?
	`

	_, err := r.DB.Exec(query, vehicle.Brand, vehicle.Model, vehicle.Year, id)

	if err != nil {
		return err
	}

	return nil
}

// DELETE A  VEHICLE
func (r *VehicleRepository) DeleteVehicle(id string) error {

	query := "DELETE FROM vehicles WHERE id = ?"

	result, err := r.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return nil
}

// GET VEHICLE DETAIL BY SERVICES
func (r *VehicleRepository) GetVehicleWithServices(id int) (models.VehicleDetail, error) {

	query := `
	SELECT 
		v.id, v.brand, v.model, v.year,
		s.id, s.description, s.cost, s.service_date
	FROM vehicles v
	LEFT JOIN services s ON v.id = s.vehicle_id
	WHERE v.id = ?
	`

	rows, err := r.DB.Query(query, id)

	if err != nil {
		return models.VehicleDetail{}, err
	}

	defer rows.Close()

	var vehicle models.VehicleDetail
	var services []models.Service

	for rows.Next() {

		var s models.Service

		err := rows.Scan(
			&vehicle.ID,
			&vehicle.Brand,
			&vehicle.Model,
			&vehicle.Year,
			&s.ID,
			&s.Description,
			&s.Cost,
			&s.ServiceDate,
		)

		if err != nil {
			return vehicle, err
		}

		// kalau service ada (tidak null)
		if s.ID != 0 {
			services = append(services, s)
		}
	}

	vehicle.Services = services

	return vehicle, nil
}