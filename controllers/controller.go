package controllers

import (
	"net/http"

	"github.com/jacky-htg/api-news/config"
	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
	"github.com/jacky-htg/api-news/repositories"
)

var userLogin models.User

func checkAuth(w http.ResponseWriter, req *http.Request, controller string, method string) bool {
	if !checkAuthToken(w, req) {
		return false
	}

	isAuth, err := repositories.AuthCheck(userLogin.Email, controller, method)
	if err != nil {
		libraries.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	if !isAuth {
		libraries.ErrorResponse(w, "Forbiden Access", http.StatusForbidden)
		return false
	}

	return true
}

func checkAuthToken(w http.ResponseWriter, req *http.Request) bool {
	if !checkApiKey(w, req) {
		return false
	}

	if len(req.Header["Token"]) == 0 {
		libraries.ErrorResponse(w, "Please suplay valid token", http.StatusBadRequest)
		return false
	}

	isTokenValid, paramEmail := libraries.ValidateToken(req.Header["Token"][0])

	if !isTokenValid {
		libraries.ErrorResponse(w, "token tidak valid", http.StatusUnauthorized)
		return false
	}

	userLogin, _ = repositories.UserGetByEmail(paramEmail)

	return true
}

func checkApiKey(w http.ResponseWriter, req *http.Request) bool {
	if len(req.Header["X-Api-Key"]) == 0 {
		libraries.ErrorResponse(w, "Please suplay valid API key", http.StatusBadRequest)
		return false
	}

	if req.Header["X-Api-Key"][0] != config.GetString("auth.apiKey") {
		libraries.ErrorResponse(w, "The API key is invalid", http.StatusUnauthorized)
		return false
	}

	return true
}

func checkError(w http.ResponseWriter, err error) bool {
	if err != nil {
		libraries.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	return true
}

func checkNull(w http.ResponseWriter, length int, err error) bool {
	if length <= 0 {
		libraries.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return false
	}

	return true
}
