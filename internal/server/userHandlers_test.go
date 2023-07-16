package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func Test_registerUserHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.POST("/rest/user").WithJSON(map[string]string{"first_name": "New", "last_name": "User"}).
		Expect().Status(http.StatusCreated).
		ContentType("application/json").
		JSON().Object().
		ContainsKey("Created user id").NotEmpty()
}

func Test_getUserHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.GET("/rest/user/ad").
		Expect().Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().Object().
		ContainsKey("error").Value("error").Equal("invalid id parameter")
}

func Test_putUserHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.PUT("/rest/user/ad").
		Expect().Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().Object().
		ContainsKey("error").Value("error").Equal("invalid id parameter")
}

func Test_deleteUserHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.DELETE("/rest/user/ad").
		Expect().Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().Object().
		ContainsKey("error").Value("error").Equal("invalid id parameter")
}
