package routes

import (
	"github.com/gorilla/mux"
	"github.com/deoliveiraromain/todo_api/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

func AddController (r *mux.Router,controller controllers.Controller) *mux.Router {
	for _, route := range controller.GetRoutes() {
		r.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return r
}

//FIXME : delete, just for test pointers
func AddTodoController (r *mux.Router,controller *controllers.TodoController) {
	for _, route := range controller.GetRoutes() {
		r.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}