package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID int64
	Name, Email string
	CreatedAt time.Time
}

type UsersModel struct {
	DB *sql.DB
}

func (model UsersModel) Insert(user *User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`

	return model.DB.QueryRow(
		query, user.Name, user.Email,
	).Scan(
		&user.ID, &user.CreatedAt,
	)
}