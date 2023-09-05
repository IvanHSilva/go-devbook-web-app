package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"web/src/config"
	"web/src/cookies"
	"web/src/requests"
	"web/src/responses"

	"github.com/gorilla/mux"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	today := time.Now()
	regDate := today.Format("02/01/2006")

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("pass"),
		"regdate":  regDate,
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}
	//fmt.Println(bytes.NewBuffer(user))

	url := fmt.Sprintf("%s/user/0", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CheckStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/unfollow", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
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

	responses.JSON(w, response.StatusCode, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/follow", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
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

	responses.JSON(w, response.StatusCode, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// today := time.Now()
	// regDate := today.Format("2006-01-02T03:04:05Z")

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		// "regdate": regDate,
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookie(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/user/%d", config.APIURL, userId)
	print(url)
	response, err := requests.DoRequestWithAuth(r, http.MethodPut, url,
		bytes.NewBuffer(user))
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

	responses.JSON(w, response.StatusCode, nil)
}
