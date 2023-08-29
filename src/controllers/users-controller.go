package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"web/src/config"
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
