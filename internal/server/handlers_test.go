package server

import (
	"net/http"
	"testing"
)

func TestGetUserHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	DataRes := struct {
		Name    string
		Surname string
	}{
		Name:    "John",
		Surname: "Silver",
	}

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody interface{}
	}{
		{"Valid ID", "/rest/user/1", http.StatusOK, DataRes},
		{"Non-existent ID", "/rest/user/2", http.StatusNotFound, "the requested resource could not be found"},
		{"Negative ID", "/rest/user/-1", http.StatusNotFound, "the requested resource could not be found"},
		{"Decimal ID", "/rest/user/1.23", http.StatusNotFound, "the requested resource could not be found"},
		{"String ID", "/rest/user/foo", http.StatusNotFound, "the requested resource could not be found"},
		{"Empty ID", "/rest/user/", http.StatusNotFound, ""},
		{"Trailing slash", "/rest/user/1/", http.StatusMovedPermanently, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			// fmt.Println(body, "---------------------------")
			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
			if body != tt.wantBody {
				t.Errorf("want body to contain %q", tt.wantBody)
			}
		})
	}

}
