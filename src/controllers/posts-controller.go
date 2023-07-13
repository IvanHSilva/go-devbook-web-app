package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requests"
	"web/src/responses"
)

func InsertPost(w http.ResponseWriter, r *http.Request) {
	today := time.Now()
	regDate := today.Format("02/01/2006")

	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
		"regdate": regDate,
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/post", config.APIURL)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url,
		bytes.NewBuffer(post))
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
