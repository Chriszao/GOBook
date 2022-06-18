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

func (repository User) FindById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM user WHERE id = ?", id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository User) FindByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, password FROM user WHERE email = ?", email,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository User) Update(id uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE user SET name = ?, nick = ?, email = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (repository User) Delete(id uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM user WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository User) FollowUser(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT ignore INTO follower (userId, followerId) VALUES (?, ?)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository User) UnFollowUser(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM follower WHERE userId = ? AND followerId = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository User) FindUserFollowers(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT user.id, user.name, user.nick, user.email, user.createdAt
		FROM user 
		INNER JOIN follower ON user.id = follower.followerId 
		WHERE follower.userId = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	var followers []models.User

	for rows.Next() {
		var follower models.User

		if err := rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository User) FindFollowing(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT user.id, user.name, user.nick, user.email, user.createdAt
		FROM user 
		INNER JOIN follower ON user.id = follower.userId 
		WHERE follower.followerId = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	var followers []models.User

	for rows.Next() {
		var follower models.User

		if err := rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository User) FindPasswordById(userId uint64) (string, error) {
	row, err := repository.db.Query(
		`SELECT password from user where id = ?`, userId,
	)

	if err != nil {
		return "", err
	}

	defer row.Close()

	var user models.User

	if row.Next() {
		if err := row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository User) UpdatePassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare(`
		UPDATE user SET password = ? WHERE id = ?
	`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
