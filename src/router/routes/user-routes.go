package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:          "/newuser",
		Method:       http.MethodGet,
		Function:     controllers.LoadUserPage,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.InsertUser,
		AuthRequired: false,
	},
	{
		URI:          "/searchuser",
		Method:       http.MethodGet,
		Function:     controllers.SearchUser,
		AuthRequired: true,
	},
	{
		URI:          "/user/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.LoadUserProfile,
		AuthRequired: true,
	},
	{
		URI:          "/user/{userId}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/user/{userId}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoggedUserProfile,
		AuthRequired: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodGet,
		Function:     controllers.EditUserPage,
		AuthRequired: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
}
