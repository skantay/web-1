package main

import (
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/skantay/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            interface{}
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for i := len(cwd); i >= 0; i-- {
		cwd = cwd[:i]
		if cwd[i-len("web-1"):i] == "web-1" {
			break
		}
	}

	pages, err := filepath.Glob(cwd + "/ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(cwd + "/ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(cwd + "/ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
