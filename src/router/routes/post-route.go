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
	{
		URI:          "/post/{postId}/edit",
		Method:       http.MethodGet,
		Function:     controllers.LoadPost,
		AuthRequired: true,
	},
	{
		URI:          "/post/{postId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/post/{postId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		AuthRequired: true,
	},
}
