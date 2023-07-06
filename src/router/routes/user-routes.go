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
}
