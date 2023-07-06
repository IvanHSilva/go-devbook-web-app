package controllers

import (
	"net/http"
	"web/src/utils"
)

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "newuser.html", nil)
}
