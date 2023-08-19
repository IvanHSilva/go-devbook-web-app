package models

import "time"

type User struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	EMail     string    `json:"email"`
	RegDate   time.Time `json:"regdate"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}
