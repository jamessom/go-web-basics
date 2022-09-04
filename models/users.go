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

func (model UsersModel) All() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`

	rows, err := model.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}