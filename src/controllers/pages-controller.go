package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
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
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError,
			responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CheckStatusCode(w, response)
		return
	}

	//fmt.Println(response.StatusCode, err)
	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}
