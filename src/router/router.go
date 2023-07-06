package router

import (
	"web/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRoutes(r)
}
