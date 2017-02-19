package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deoliveiraromain/task-rest-api/db"
	"github.com/deoliveiraromain/task-rest-api/handlers"
	"github.com/deoliveiraromain/task-rest-api/routes"
	"gopkg.in/mgo.v2"
)

func main() {
	//get config
	conf, err := configuration.GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}
	// Connect to our local mongo

	s, err := mgo.Dial("mongodb://" + conf.MongoHost)
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

	portStr := fmt.Sprintf(":%d", conf.Port)
	log.Printf("Serving on %s", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}

// ServeHTTP is the http.Handler interface implementation
func serveWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome on TODO list API written in GO.!\n")
}
