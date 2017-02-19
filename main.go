package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/deoliveiraromain/task-rest-api/configuration"
	"github.com/deoliveiraromain/task-rest-api/db"
	"github.com/deoliveiraromain/task-rest-api/handlers"
	"github.com/deoliveiraromain/task-rest-api/routes"
)

func main() {
	//get config
	conf, err := configuration.GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}
	// Connect to our local mongo
	log.Printf("Config port %s", conf.Port)
	log.Printf("Config MongoHost %s", conf.MongoHost)
	maxWait := time.Duration(5 * time.Second)
	s, err := mgo.DialWithTimeout("mongodb://"+conf.MongoHost+":27017", maxWait)
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

	portStr := fmt.Sprintf(":%s", conf.Port)
	log.Printf("Serving on %s", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}

// ServeHTTP is the http.Handler interface implementation
func serveWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome on TODO list API written in GO.!\n")
}
