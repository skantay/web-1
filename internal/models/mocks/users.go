package mocks

import (
	"database/sql"
	"github.com/skantay/snippetbox/internal/models"
)

type UserModel struct {
	*sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	if email == "bob@example.com" {
		return nil
	}

	return models.ErrDuplicateEmail
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "mock@mock.com" && password == "123" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	if id == 1 {
		return true, nil
	}

	return false, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

func (m *UserModel) PasswordUpdate(oldPass, newPass string, id int) error {
	return nil
}
