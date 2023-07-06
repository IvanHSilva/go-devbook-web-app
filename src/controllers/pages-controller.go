package controllers

import (
	"net/http"
	"web/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "newuser.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
