package repository

import (
	"database/sql"
	"devbook-api/internal/model"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository *Users) Create(user *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, nickname, email, password) VALUES ($1, $2, $3, $4) RETURNING id, name, nickname, email, created_at"

	var userCreated model.User = *user

	err := repository.db.QueryRow(
		query,
		user.Name,
		user.Nickname,
		user.Email,
		user.Password,
	).Scan(&userCreated.ID,
		&userCreated.Name,
		&userCreated.Nickname,
		&userCreated.Email,
		&userCreated.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &userCreated, nil
}
