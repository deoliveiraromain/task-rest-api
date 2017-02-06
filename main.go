package main

import (
	"log"
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/deoliveiraromain/todo_api/routes"
	"fmt"
	"github.com/deoliveiraromain/todo_api/handlers"
	"gopkg.in/mgo.v2"
)

func main() {

	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	defer s.Close()

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	con := db.NewMongo(s, "todos", "todo")

	router := routes.NewRouter()
	router.HandleFunc("/", serveWelcome)

	tc := handlers.NewTodoController(con)
	tc.Register(router)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// ServeHTTP is the http.Handler interface implementation
func serveWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome on TODO list API written in GO.!\n")
}