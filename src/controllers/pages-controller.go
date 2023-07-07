package controllers

import (
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requests"
	"web/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "newuser.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/posts", config.APIURL)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
