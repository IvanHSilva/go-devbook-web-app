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
}
