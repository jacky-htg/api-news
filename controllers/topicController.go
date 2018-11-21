package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
	"github.com/jacky-htg/api-news/repositories"
)

func TopicListHandler(w http.ResponseWriter, r *http.Request) {
	if !checkApiKey(w, r) {
		return
	}

	params := r.URL.Query()
	page := "1"
	if len(params["page"]) > 0 {
		page = params["page"][0]
	}

	limit := "50"
	if len(params["limit"]) > 0 {
		limit = params["limit"][0]
	}

	search := ""
	if len(params["search"]) > 0 {
		search = params["search"][0]
	}

	sortby := "id"
	if len(params["sortby"]) > 0 {
		sortby = params["sortby"][0]
	}

	order := "desc"
	if len(params["order"]) > 0 {
		order = params["order"][0]
	}

	if there, _ := libraries.InArray(order, []string{"ASC", "DESC", "asc", "desc"}); !there {
		order = "desc"
	}

	var param = map[string]string{"page": page, "limit": limit, "search": search, "sortby": sortby, "order": order}

	topics, err := repositories.TopicGetList(param)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, topics, http.StatusOK)
}

func TopicGetHandler(w http.ResponseWriter, r *http.Request) {
	if !checkApiKey(w, r) {
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if !checkError(w, err) {
		return
	}

	topic, err := repositories.TopicGet(uint(id))
	if !checkError(w, err) {
		return
	}

	if topic.ID < 1 {
		libraries.ErrorResponse(w, "Topic not found", http.StatusNotFound)
		return
	}

	libraries.SetData(w, topic, http.StatusOK)
}

func TopicCreateHandler(w http.ResponseWriter, r *http.Request) {
	method := "store"
	controller := "topic"
	if !checkAuth(w, r, controller, method) {
		return
	}

	var oR models.Topic

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&oR)

	if !checkError(w, err) {
		return
	}

	defer r.Body.Close()

	if len(oR.Title) <= 0 || oR.Title == "" {
		libraries.ErrorResponse(w, "Please insert valid title!", http.StatusBadRequest)
		return
	}

	isTopicExist, err := repositories.TopicIsExist(oR.Title, 0)
	if !checkError(w, err) {
		return
	}

	if isTopicExist {
		libraries.ErrorResponse(w, "Topic already exist!", http.StatusInternalServerError)
		return
	}

	oR = models.TopicValidate(oR)
	topic, err := repositories.TopicStore(oR)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, topic, http.StatusCreated)
}

func TopicUpdateHandler(w http.ResponseWriter, r *http.Request) {
	method := "update"
	controller := "topic"
	if !checkAuth(w, r, controller, method) {
		return
	}

	var oR models.Topic

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&oR)
	if !checkError(w, err) {
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if !checkError(w, err) {
		return
	}
	oR.ID = uint(id)

	defer r.Body.Close()

	topicExisting, err := repositories.TopicGet(oR.ID)
	if !checkError(w, err) {
		return
	}

	if topicExisting.ID < 1 {
		libraries.ErrorResponse(w, "Topic not found", http.StatusNotFound)
		return
	}

	if len(oR.Title) <= 0 && oR.Title == "" {
		libraries.ErrorResponse(w, "Please insert valid title!", http.StatusBadRequest)
		return
	}

	isTopicExist, err := repositories.TopicIsExist(oR.Title, oR.ID)
	if !checkError(w, err) {
		return
	}

	if isTopicExist {
		libraries.ErrorResponse(w, "Topic already exist!", http.StatusInternalServerError)
		return
	}

	oR = models.TopicValidate(oR)
	topic, err := repositories.TopicUpdate(oR)
	if !checkError(w, err) {
		return
	}

	topic, err = repositories.TopicGet(topic.ID)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, topic, http.StatusOK)
}

func TopicDeleteHandler(w http.ResponseWriter, r *http.Request) {
	method := "destroy"
	controller := "topic"
	if !checkAuth(w, r, controller, method) {
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.Atoi(id)
	if !checkError(w, err) {
		return
	}

	topicExisting, err := repositories.TopicGet(uint(ID))
	if !checkError(w, err) {
		return
	}

	if topicExisting.ID < 1 {
		libraries.ErrorResponse(w, "Topic not found", http.StatusNotFound)
		return
	}

	_, err = repositories.TopicDestroy(uint(ID))
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, nil, http.StatusNoContent)
}
