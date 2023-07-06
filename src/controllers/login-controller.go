package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	response, err := http.Post("http://localhost:5000/login", "application/json",
		bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	//defer response.Body.Close()

	token, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(token))

	// if response.StatusCode >= 400 {
	// 	responses.CheckStatusCode(w, response)
	// 	return
	// }

	// responses.JSON(w, response.StatusCode, nil)
}
