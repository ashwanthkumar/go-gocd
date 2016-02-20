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

const testUsername = "admin"
const testPassword = "badger"

func newTestClient(t *testing.T) *Client {
	once.Do(func() {
		testServerHandler := http.NewServeMux()
		// TODO - Add more handlers here as we implement more functionalities of the client
		testServerHandler.HandleFunc("/go/api/agents", serveFileAsJSON(t, "test-fixtures/get_all_agents.json"))
		testServerHandler.HandleFunc("/go/api/jobs/scheduled.xml", serveFileAsXML(t, "test-fixtures/get_scheduled_jobs.xml"))
		testServerHandler.HandleFunc("/go/api/jobs/pipeline/stage/job/history/0", serveFileAsJSON(t, "test-fixtures/get_job_history.json"))
		apiServer = httptest.NewServer(testServerHandler)
	})

	return New(apiServer.URL, testUsername, testPassword)
}

func serveFileAsJSON(t *testing.T, filepath string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		AcceptHeaderCheck(t, request)
		BasicAuthCheck(t, request)

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Fatal(err)
		}
		writer.Header().Add("Content-Type", "application/vnd.go.cd.v2+json; charset=utf-8")
		writer.Write(contents)
	}
}
func serveFileAsXML(t *testing.T, filepath string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		BasicAuthCheck(t, request)

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Fatal(err)
		}
		writer.Header().Add("Content-Type", "application/xml; charset=utf-8")
		writer.Write(contents)
	}
}

func AcceptHeaderCheck(t *testing.T, request *http.Request) {
	// Accept Header check
	acceptHeader := request.Header.Get("Accept")
	if acceptHeader != "application/vnd.go.cd.v2+json" {
		t.Fatal("We did not recieve Accept: application/vnd.go.cd.v2+json header in the request")
	}
}

func BasicAuthCheck(t *testing.T, request *http.Request) {
	// BasicAuth check
	username, password, _ := request.BasicAuth()
	if username != testUsername && password != testPassword {
		t.Fatal("Invalid username / password combination")
	}
}
