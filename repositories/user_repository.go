// SESUDAH (fix)
package repositories

import (
	"database/sql"
	"errors"
	"garage-api/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT id, username, password, role FROM users WHERE username = ?"

	err := r.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}