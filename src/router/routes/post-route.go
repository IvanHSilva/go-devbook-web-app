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
		AuthRequired: true,
	},
	{
		URI:          "/post/{postId}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/post/{postId}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePost,
		AuthRequired: true,
	},
}
