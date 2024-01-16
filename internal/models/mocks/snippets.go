package mocks

import (
	"time"

	"github.com/skantay/snippetbox/internal/models"
)


var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "MOCK TITLE",
	Content: "MOCK CONTENT",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 1, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	if id == 1 {
		return mockSnippet, nil
	}

	return nil, models.ErrNoRecord
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}

