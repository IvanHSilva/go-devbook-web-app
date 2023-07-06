package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
		log.Fatal(err)
	}
	//fmt.Println(bytes.NewBuffer(user))

	response, err := http.Post("http://localhost:5000/user/0", "application/json",
		bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}
