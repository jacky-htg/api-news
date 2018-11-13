package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/repositories"
	"golang.org/x/crypto/bcrypt"
)

type authLogin struct {
	Email    string
	Password string
}

type token struct {
	Token string
}

func AuthGetTokenHandler(w http.ResponseWriter, req *http.Request) {
	if !checkApiKey(w, req) {
		return
	}

	var login authLogin

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&login)

	if !checkError(w, err) {
		return
	}
	defer req.Body.Close()

	if login.Email == "" || login.Password == "" {
		libraries.ErrorResponse(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	userToken, err := repositories.UserGetByEmail(login.Email)
	if !checkError(w, err) {
		return
	}

	if userToken.ID <= 0 {
		libraries.ErrorResponse(w, "user not found", http.StatusNotFound)
		return
	}

	if !userToken.IsActive {
		libraries.ErrorResponse(w, "user inactivated", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userToken.Password), []byte(login.Password))
	if err != nil {
		libraries.ErrorResponse(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	tokenString, err := libraries.ClaimToken(login.Email)
	if !checkError(w, err) {
		return
	}

	libraries.SetData(w, token{Token: tokenString}, http.StatusOK)
}
