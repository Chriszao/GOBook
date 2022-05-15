package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (repository User) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO user (name, nick, email, password) VALUES (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}

func (repository User) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // Output: %nameOrNick%

	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM user WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
