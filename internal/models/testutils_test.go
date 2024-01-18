package models

import (
	"database/sql"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for i := len(cwd); i >= 0; i-- {
		cwd = cwd[:i]
		if cwd[i-len("web-1"):i] == "web-1" {
			break
		}
	}

	script, err := os.ReadFile(cwd + "/internal/models/testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(string(script)); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile(cwd + "/internal/models/testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		if _, err := db.Exec(string(script)); err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	return db
}
