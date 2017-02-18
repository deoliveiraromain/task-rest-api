package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deoliveiraromain/task-rest-api/db"
	"github.com/deoliveiraromain/task-rest-api/handlers"
	"github.com/deoliveiraromain/task-rest-api/routes"
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
	con := db.NewMongo(s, "todos")

	router := routes.NewRouter()
	router.HandleFunc("/", serveWelcome)

	tc := handlers.NewTaskController(con)
	tc.Register(router)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// ServeHTTP is the http.Handler interface implementation
func serveWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome on TODO list API written in GO.!\n")
}
