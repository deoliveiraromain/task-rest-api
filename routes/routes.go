package routes

import (
	"github.com/gorilla/mux"
	"github.com/deoliveiraromain/todo_api/controllers"
)

func NewRouter(controllers []controllers.Controller) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, controller := range controllers {
		for _, route := range controller.GetRoutes() {
			router.
			Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
	return router
}