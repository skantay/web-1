package mocks

import (
	"github.com/skantay/snippetbox/internal/models"
)

type UserModel struct{}

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
