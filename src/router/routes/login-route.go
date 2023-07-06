package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.Login,
		AuthRequired: false,
	},
}
