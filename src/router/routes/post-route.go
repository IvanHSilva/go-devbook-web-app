package routes

import (
	"net/http"
	"web/src/controllers"
)

var postRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.InsertPost,
		AuthRequired: false,
	},
}