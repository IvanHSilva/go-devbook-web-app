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
}
