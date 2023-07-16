package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func Test_valHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.GET("/rest/counter/val/").Expect().Status(http.StatusOK).ContentType("application/json").JSON().Object().
		ContainsKey("counter").Value("counter")
}

func Test_addHandler(t *testing.T) {
	app := newTestApplication(t)

	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.POST("/rest/counter/add/5").
		Expect().
		Status(http.StatusOK).
		ContentType("application/json").JSON().
		Object().ContainsKey("counter").Value("counter")
}

func Test_subHandler(t *testing.T) {
	app := newTestApplication(t)

	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.POST("/rest/counter/sub/5").
		Expect().
		Status(http.StatusOK).
		ContentType("application/json").JSON().
		Object().ContainsKey("counter").Value("counter")
}
