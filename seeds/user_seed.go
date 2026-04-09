package seeds

import (
	"database/sql"
	"fmt"
	"garage-api/helpers"
)

func SeedUser(db *sql.DB) {

	username := "user1"
	password := "123456"

	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		panic(err)
	}

	query := `
	INSERT INTO users (username, password)
	VALUES (?, ?)
	`

	_, err = db.Exec(query, username, hashedPassword)

	if err != nil {
		fmt.Println("Seeder error:", err)
		return
	}

	fmt.Println("User seeded successfully")
}