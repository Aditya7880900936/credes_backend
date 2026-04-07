package repository

import (
	"github.com/Aditya7880900936/credes-backend/internal/db"
	"github.com/Aditya7880900936/credes-backend/internal/models"
)

func CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (email, password, full_name, role)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return db.DB.QueryRowx(query,
		user.Email,
		user.Password,
		user.FullName,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt)
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email=$1`

	err := db.DB.Get(&user, query, email)
	return &user, err
}
