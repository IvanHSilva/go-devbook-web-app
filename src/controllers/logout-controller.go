package controllers

import (
	"net/http"
	"web/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	cookies.CleanCookie(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
