package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/skantay/snippetbox/internal/assert"
)

func TestSnippetCreate(t *testing.T) {

	tests := []struct {
		name          string
		wantCode      int
		wantHeader    string
		wantBody      string
		authenticated bool
	}{
		{
			name:          "Unauthenticated",
			wantCode:      303,
			wantHeader:    "/user/login",
			authenticated: false,
		},
		{
			name:          "Authenticated",
			wantCode:      200,
			wantBody:      `<form action='/snippet/create' method='POST'>`,
			authenticated: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			app := newTestApplication(t)
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			if testCase.authenticated {
				_, _, body := ts.get(t, "/user/login")
				validCSRFToken := extractCSRFToken(t, body)

				form := url.Values{}
				form.Add("email", "mock@mock.com")
				form.Add("password", "123")
				form.Add("csrf_token", validCSRFToken)
				ts.postForm(t, "/user/login", form)

			}

			code, header, body := ts.get(t, "/snippet/create")

			assert.Equal(t, code, testCase.wantCode)

			if !testCase.authenticated {
				assert.Equal(t, header.Get("Location"), testCase.wantHeader)
			} else {
				assert.StringContains(t, body, testCase.wantBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Bob"
		validPassword = "validPa$$word"
		validEmail    = "bob@example.com"
		formTag       = "<form action=\"/user/signup\" method=\"POST\" novalidate>"
	)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantFormTag  string
	}{
		{
			name:         "Valid submission",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "Empty name",
			userName:     "",
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty email",
			userName:     validName,
			userEmail:    "",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
		},
		{
			name:         "Duplicate email",
			userName:     validName,
			userEmail:    "123@123.123",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", testCase.userName)
			form.Add("email", testCase.userEmail)
			form.Add("password", testCase.userPassword)
			form.Add("csrf_token", testCase.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, testCase.wantCode)

			if testCase.wantFormTag != "" {
				assert.StringContains(t, body, testCase.wantFormTag)
			}
		})
	}
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	testCases := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "MOCK CONTENT",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/asdf",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			code, _, body := ts.get(t, testCase.urlPath)

			assert.Equal(t, code, testCase.wantCode)
			if testCase.wantBody != "" {
				assert.StringContains(t, body, testCase.wantBody)
			}
		})
	}
}

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
