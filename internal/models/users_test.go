package models

import (
	"github.com/skantay/snippetbox/internal/assert"
	"testing"
)

func TestUserModelExists(t *testing.T) {
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "two ID",
			userID: 2,
			want:   false,
		},
	}

	for _, testCase := range tests {

		t.Run(testCase.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			exists, err := m.Exists(testCase.userID)

			assert.Equal(t, exists, testCase.want)
			assert.NillError(t, err)
		})
	}
}
