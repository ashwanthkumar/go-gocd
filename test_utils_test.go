package gocd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var once sync.Once
var apiServer *httptest.Server

func newTestClient(t *testing.T) *Client {
	once.Do(func() {
		testServerHandler := http.NewServeMux()
		// TODO - Add more handlers here as we implement more functionalities of the client
		testServerHandler.HandleFunc("/go/api/agents", serveFileAsJSON(t, "test-fixtures/get_all_agents.json"))
		apiServer = httptest.NewServer(testServerHandler)
	})

	return New(apiServer.URL, "foo", "bar")
}

func serveFileAsJSON(t *testing.T, filepath string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, reader *http.Request) {
		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Fatal(err)
		}
		writer.Header().Add("Content-Type", "application/vnd.go.cd.v1+json; charset=utf-8")
		writer.Write(contents)
	}
}
