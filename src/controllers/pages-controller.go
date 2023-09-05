package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.ReadCookie(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

func LoadPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/post/%d", config.APIURL, postId)

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

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "editpost.html", post)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {

	user := strings.ToLower(r.URL.Query().Get("user"))
	print(user)
	url := fmt.Sprintf("%s/user?user=%s", config.APIURL, user)
	print(url)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == loggedUserId {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	user, err := models.SearchAllUserData(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError,
			responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserId uint64
	}{
		User:         user,
		LoggedUserId: loggedUserId,
	})
}

func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchAllUserData(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}
