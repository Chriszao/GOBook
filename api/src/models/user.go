package models

import (
	"api/src/providers"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(operation string) error {
	if err := user.validate(operation); err != nil {
		return err
	}

	if err := user.formatData(operation); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(operation string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email format is invalid")
	}

	if operation == "new" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) formatData(operation string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if operation == "new" {
		hashedPassword, err := providers.Hash(user.Password)

		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}
	return nil
}
