package routes

import (
	"net/http"
	"web/src/controllers"
)

var homeRoute = Route{
	URI:          "/home",
	Method:       http.MethodGet,
	Function:     controllers.LoadHomePage,
	AuthRequired: true,
}
