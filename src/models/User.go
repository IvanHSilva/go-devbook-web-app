package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requests"
)

type User struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	EMail     string    `json:"email"`
	RegDate   time.Time `json:"regdate"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func SearchAllUserData(userId uint64, r *http.Request) (User, error) {
	// channels
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowing := make(chan []User)
	channelPosts := make(chan []Post)

	// functions
	go SearchUserData(channelUser, userId, r)
	go SearchFollowers(channelFollowers, userId, r)
	go SearchFollowing(channelFollowing, userId, r)
	go SearchPosts(channelPosts, userId, r)

	// variables
	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	// check channels
	for i := 1; i <= 4; i++ {
		select {
		case userChannel := <-channelUser:
			if userChannel.Id == 0 {
				return User{}, errors.New("erro ao buscar usuário")
			}
			user = userChannel
		case followersChannel := <-channelFollowers:
			if followersChannel == nil {
				return User{}, errors.New("erro ao buscar seguidores")
			}
			followers = followersChannel
		case followingChannel := <-channelFollowing:
			if followingChannel == nil {
				return User{}, errors.New("erro ao buscar usuários seguindo")
			}
			followers = followingChannel
		case postsChannel := <-channelPosts:
			if postsChannel == nil {
				return User{}, errors.New("erro ao buscar postagens")
			}
			posts = postsChannel
		}
	}

	// return data
	user.Followers = followers
	//if following != nil {
	user.Following = following
	//}
	user.Posts = posts
	return user, nil
}

func SearchUserData(channel chan<- User, userId uint64, r *http.Request) {

	url := fmt.Sprintf("%s/user/%d", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func SearchFollowers(channel chan<- []User, userId uint64, r *http.Request) {

	url := fmt.Sprintf("%s/user/%d/followers", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func SearchFollowing(channel chan<- []User, userId uint64, r *http.Request) {

	url := fmt.Sprintf("%s/user/%d/followers", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func SearchPosts(channel chan<- []Post, userId uint64, r *http.Request) {

	url := fmt.Sprintf("%s/user/%d/post", config.APIURL, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
