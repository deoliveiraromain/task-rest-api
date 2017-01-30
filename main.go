package main

import (
	"log"
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/deoliveiraromain/todo_api/routes"
	"github.com/deoliveiraromain/todo_api/controllers"
)

func main() {

	var database db.DB
	database = db.NewMem()

	tc := controllers.NewTodoController(database)

	arr := [1]controllers.Controller{tc}
	router := routes.NewRouter(arr)
	log.Fatal(http.ListenAndServe(":8080", router))
}
