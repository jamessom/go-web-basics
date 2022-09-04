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

func (model UsersModel) All(filter Filter) ([]User, Metadata, error) {
	var users []User
	query := `SELECT
							COUNT(*) OVER(),
							id, name, email, created_at
						FROM
							users
						LIMIT
							$1
						OFFSET
							$2`

	rows, err := model.DB.Query(query, filter.Limit(), filter.Offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	var totalRec int
	for rows.Next() {
		var user User
		err = rows.Scan(&totalRec, &user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, Metadata{}, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	return users, ComputeMetadata(totalRec, filter.Page, filter.PageSize), nil
}