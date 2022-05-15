package repositories

import (
	"api/src/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (userRepository User) Create(user models.User) (uint64, error) {
	return 0, nil
}
