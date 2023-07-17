package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"Alikhan.Aitbayev/Data"
	"Alikhan.Aitbayev/Data/mock"
)

func newTestApplication(t *testing.T) *Application {

	return &Application{
		ErrorLog: log.New(ioutil.Discard, "", 0),
		InfoLog:  log.New(ioutil.Discard, "", 0),
		models: Data.Models{
			Users: mock.UserModel{},
		},
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, interface{}) {
	type MyData struct {
		Error string `json:"error"`
	}
	myData := MyData{}

	Data := map[string]map[string]string{}

	DataRes := struct {
		Name    string
		Surname string
	}{}
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()
	flag := 1
	if rs.Status == "404 Not Found" && rs.ContentLength != 19 {
		err = json.NewDecoder(rs.Body).Decode(&myData)
		// fmt.Println(myData, "---------------")
	} else if rs.Status == "200 OK" {
		flag = 0
		err = json.NewDecoder(rs.Body).Decode(&Data)
		// fmt.Println(Data, "-----------------")
	}

	if err != nil {
		t.Fatal(err)
	}
	if flag == 0 {
		DataRes.Name = Data["user"]["first_name"]
		DataRes.Surname = Data["user"]["last_name"]
		return rs.StatusCode, rs.Header, DataRes
	} else {
		return rs.StatusCode, rs.Header, myData.Error
	}

}
