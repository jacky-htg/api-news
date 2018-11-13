package main_test

import (
	//"log"
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-news/config"
	"github.com/jacky-htg/api-news/controllers"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestTopicList(t *testing.T) {

	req, _ := http.NewRequest("GET", "/topics", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/topics", controllers.TopicListHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"code":200,"data":[{"id":3,"title":"Jakarta","created_at":"2018-11-11T17:39:04Z","updated_at":"2018-11-11T17:39:04Z"},{"id":2,"title":"Nasional","created_at":"2018-11-11T15:38:34Z","updated_at":"2018-11-11T15:38:34Z"},{"id":1,"title":"Pilpres","created_at":"2018-11-11T15:37:41Z","updated_at":"2018-11-11T15:37:41Z"}]}
`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTopicGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/topics/1", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/topics/{id}", controllers.TopicGetHandler).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"code":200,"data":{"id":1,"title":"Pilpres","created_at":"2018-11-11T15:37:41Z","updated_at":"2018-11-11T15:37:41Z"}}
`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
