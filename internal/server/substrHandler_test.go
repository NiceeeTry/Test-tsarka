package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func Test_substrHandler(t *testing.T) {
	app := newTestApplication(t)
	handler := app.routes()
	server := httptest.NewServer(handler)
	defer server.Close()

	api := httpexpect.New(t, "http://127.0.0.1:8080")

	api.POST("/rest/substr/find").WithJSON(map[string]string{"text": "asdnsasdf"}).
		Expect().Status(http.StatusOK).
		ContentType("application/json").
		JSON().Object().
		ContainsKey("Result").Value("Result").IsEqual("asdn")
}
