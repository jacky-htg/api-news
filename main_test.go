package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-news/config"
	"github.com/jacky-htg/api-news/models"
	"github.com/jacky-htg/api-news/repositories"
	"github.com/jacky-htg/api-news/routing"
)

type result_token struct {
	Code    int    `json:"code"`
	Data    token  `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type result_json struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type token struct {
	Token string
}

var router *mux.Router
var tokenStr string

func TestMain(m *testing.M) {
	router = routing.NewRouter()
	code := m.Run()
	os.Exit(code)
}

func TestGetToken(t *testing.T) {
	var jsonStr = []byte(`{"email":"editor@gmail.com","password":"1234"}`)
	req, _ := http.NewRequest("POST", "/get-token", bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var m result_token
	err := json.NewDecoder(rr.Body).Decode(&m)
	if err != nil {
		t.Error(err.Error())
	}
	tokenStr = m.Data.Token
}

func TestTopicList(t *testing.T) {

	req, _ := http.NewRequest("GET", "/topics?limit=1&order=desc", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	topic, err := repositories.TopicFindLast()
	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	var topics []models.Topic
	topics = append(topics, topic)

	result.Code = http.StatusOK
	result.Data = topics

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/topics/1", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	topic, err := repositories.TopicFindFirst()
	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusOK
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicCreate(t *testing.T) {

	req, _ := http.NewRequest("POST", "/topics", bytes.NewBuffer([]byte(`{"title": "Topik Baru"}`)))
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusCreated
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicUpdate(t *testing.T) {

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	req, _ := http.NewRequest("PUT", "/topics/"+strconv.Itoa(int(topic.ID)), bytes.NewBuffer([]byte(`{"title": "Topik Baru Edit"}`)))
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	topic, err = repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusOK
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicDestroy(t *testing.T) {

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	req, _ := http.NewRequest("DELETE", "/topics/"+strconv.Itoa(int(topic.ID)), nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestNewsList(t *testing.T) {

	req, _ := http.NewRequest("GET", "/news?limit=1&order=desc", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	last, err := repositories.NewsFindLast()
	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	var news []models.News
	news = append(news, last)

	result.Code = http.StatusOK
	result.Data = news

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

/*func TestTopicGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/topics/1", nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	topic, err := repositories.TopicFindFirst()
	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusOK
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicCreate(t *testing.T) {

	req, _ := http.NewRequest("POST", "/topics", bytes.NewBuffer([]byte(`{"title": "Topik Baru"}`)))
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusCreated
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicUpdate(t *testing.T) {

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	req, _ := http.NewRequest("PUT", "/topics/"+strconv.Itoa(int(topic.ID)), bytes.NewBuffer([]byte(`{"title": "Topik Baru Edit"}`)))
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	topic, err = repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	var result result_json
	result.Code = http.StatusOK
	result.Data = topic

	res, err := json.Marshal(result)
	if err != nil {
		t.Error(err.Error())
	}

	body := rr.Body.String()
	body = body[:(len(body) - 1)]
	if body != string(res) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, string(res))
	}
}

func TestTopicDestroy(t *testing.T) {

	topic, err := repositories.TopicFindLast()

	if err != nil {
		t.Error(err.Error())
	}

	req, _ := http.NewRequest("DELETE", "/topics/"+strconv.Itoa(int(topic.ID)), nil)
	req.Header.Set("X-Api-Key", config.GetString("auth.apiKey"))
	req.Header.Set("Token", tokenStr)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}*/
