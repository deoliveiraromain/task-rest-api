package main

import (
	"log"
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/deoliveiraromain/todo_api/routes"
	"fmt"
	"github.com/deoliveiraromain/todo_api/handlers"
	"github.com/deoliveiraromain/todo_api/models"
)

var todos models.Todos

func init() {
	todos = models.Todos{
		models.Todo{
			1,
			"Write presentation",
			false,
			//"0001-01-01T00:00:00Z",
		}, models.Todo{
			2,
			"Host meetup",
			false,
			//"0001-01-01T00:00:00Z",
		}, models.Todo{
			3,
			"New Todo",
			false,
			//"0001-01-01T00:00:00Z",
		},
	}
}
func main() {

	var database db.DB
	database = db.NewMem()

	router := routes.NewRouter()
	router.HandleFunc("/", serveWelcome)

	tc := handlers.NewTodoController(database)
	tc.Register(router)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// ServeHTTP is the http.Handler interface implementation
func serveWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome on TODO list API written in GO.!\n")
}