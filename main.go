package main

import (
	"log"
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/deoliveiraromain/todo_api/routes"
	"github.com/deoliveiraromain/todo_api/controllers"
	"fmt"
)

func main() {

	var database db.DB
	database = db.NewMem()

	tc := controllers.NewTodoController(database)

	router := routes.NewRouter()
	router.HandleFunc("/", ServeWelcome)
	router.Methods("GET").
		Path("/caca").
		Name("getTodo").
		HandlerFunc(Todo)
	router = routes.AddController(router, tc)
	http.Handle()
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
// ServeHTTP is the http.Handler interface implementation
func ServeWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Todo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome TODO!\n")
}