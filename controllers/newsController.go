package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
	"github.com/jacky-htg/api-news/repositories"
)

func NewsGetHandler(w http.ResponseWriter, r *http.Request) {
	if !checkApiKey(w, r) {
		return
	}

	params := mux.Vars(r)
	paramID, err := strconv.Atoi(params["id"])
	if !checkError(w, err) {
		return
	}

	news, err := repositories.NewsGet(uint(paramID))
	if !checkError(w, err) {
		return
	}

	if news.ID < 1 || news.Status == "X" {
		libraries.ErrorResponse(w, "News not found", http.StatusNotFound)
		return
	}

	libraries.SetData(w, news, http.StatusOK)
}

func NewsListHandler(w http.ResponseWriter, r *http.Request) {
	if !checkApiKey(w, r) {
		return
	}

	params := r.URL.Query()
	status := ""
	if len(params["status"]) > 0 {
		status = params["status"][0]
	}

	if there, _ := libraries.InArray(status, []string{"D", "X", "P"}); !there {
		status = ""
	}

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

	topic := ""
	if len(params["topic"]) > 0 {
		topic = params["topic"][0]
	}

	var param = map[string]string{"status": status, "page": page, "limit": limit, "search": search, "topic": topic, "sortby": sortby, "order": order}

	news, err := repositories.NewsList(param)
	if !checkError(w, err) {
		return
	}

	if !checkNull(w, len(news), errors.New("News not found")) {
		return
	}

	libraries.SetData(w, news, http.StatusOK)
}

func NewsCreateHandler(w http.ResponseWriter, r *http.Request) {
	method := "store"
	controller := "news"
	if !checkAuth(w, r, controller, method) {
		return
	}

	var oR models.News

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&oR)

	if !checkError(w, err) {
		return
	}

	defer r.Body.Close()
	oR.Writer.ID = userLogin.ID

	if len(oR.Title) <= 0 || oR.Title == "" {
		libraries.ErrorResponse(w, "Please insert valid title!", http.StatusBadRequest)
		return
	}

	isNewsExists, err := repositories.NewsIsExists(oR.Title, 0)
	if !checkError(w, err) {
		return
	}

	if isNewsExists {
		libraries.ErrorResponse(w, "News with this title is already exists!, please change with others", http.StatusInternalServerError)
		return
	}

	oR.Slug = slug.Make(oR.Title)

	if len(oR.Content) <= 0 {
		oR.Content = ""
	}

	news, err := repositories.NewsStore(oR)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, news, http.StatusOK)
}

func NewsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	method := "store"
	controller := "news"
	if !checkAuth(w, r, controller, method) {
		return
	}

	var oR models.News

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

	newsExisting, err := repositories.NewsGet(oR.ID)
	if !checkError(w, err) {
		return
	}

	if newsExisting.Writer.ID != userLogin.ID {
		if !checkAuth(w, r, controller, "update") {
			return
		}
	}

	if newsExisting.ID < 1 || newsExisting.Status == "X" {
		libraries.ErrorResponse(w, "News not found", http.StatusNotFound)
		return
	}

	if newsExisting.Status == "P" {
		libraries.ErrorResponse(w, "News has published. Can not to be updated", http.StatusForbidden)
		return
	}

	oR.Editor = userLogin

	if len(oR.Title) > 0 {
		if oR.Title == "" {
			libraries.ErrorResponse(w, "Please insert valid title!", http.StatusBadRequest)
			return
		}

		isNewsExist, err := repositories.NewsIsExists(oR.Title, oR.ID)
		if !checkError(w, err) {
			return
		}

		if isNewsExist {
			libraries.ErrorResponse(w, "News already exist! Please change the title.", http.StatusInternalServerError)
			return
		}

		oR.Slug = slug.Make(oR.Title)
	}

	news, err := repositories.NewsUpdate(oR)
	if !checkError(w, err) {
		return
	}

	news, err = repositories.NewsGet(news.ID)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, news, http.StatusOK)
}

func NewsPublishHandler(w http.ResponseWriter, r *http.Request) {
	method := "publish"
	controller := "news"
	if !checkAuth(w, r, controller, method) {
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if !checkError(w, err) {
		return
	}

	newsExisting, err := repositories.NewsGet(uint(id))
	if !checkError(w, err) {
		return
	}

	if newsExisting.ID < 1 || newsExisting.Status == "X" {
		libraries.ErrorResponse(w, "News not found", http.StatusNotFound)
		return
	}

	if newsExisting.Status == "P" {
		libraries.ErrorResponse(w, "News has published", http.StatusForbidden)
		return
	}

	if newsExisting.PublishDate.Format(time.RFC822) == "01 Jan 01 00:00 UTC" {
		libraries.ErrorResponse(w, "Publish date can not null", http.StatusForbidden)
		return
	}

	newsExisting.Editor = userLogin

	news, err := repositories.NewsPublish(newsExisting)
	if !checkError(w, err) {
		return
	}

	news, err = repositories.NewsGet(news.ID)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, news, http.StatusOK)
}

func NewsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	method := "destroy"
	controller := "news"
	if !checkAuth(w, r, controller, method) {
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if !checkError(w, err) {
		return
	}

	newsExisting, err := repositories.NewsGet(uint(id))
	if !checkError(w, err) {
		return
	}

	if newsExisting.ID < 1 || newsExisting.Status == "X" {
		libraries.ErrorResponse(w, "News not found", http.StatusNotFound)
		return
	}

	_, err = repositories.NewsDestroy(newsExisting)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, nil, http.StatusNoContent)
}
