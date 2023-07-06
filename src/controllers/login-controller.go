package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/models"
	"web/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("pass"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
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

	var dataAuth models.DataAuth
	if err = json.NewDecoder(response.Body).Decode(&dataAuth); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
