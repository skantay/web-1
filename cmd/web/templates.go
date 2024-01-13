package main

import "github.com/skantay/snippetbox/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
