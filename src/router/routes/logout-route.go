package routes

import (
	"net/http"
	"web/src/controllers"
)

var logoutRoute = Route{
	URI:          "/logout",
	Method:       http.MethodGet,
	Function:     controllers.Logout,
	AuthRequired: true,
}
