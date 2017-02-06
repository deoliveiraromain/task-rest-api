package handlers

import (
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/deoliveiraromain/todo_api/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

type TodoController struct {
	db *db.Mongo
}

func NewTodoController(db *db.Mongo) *TodoController {
	tc := &TodoController{db :db}
	return tc
}
func (tc *TodoController) Register(router *mux.Router) {
	router.HandleFunc("/todos", tc.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{Name}", tc.GetTodoByName).Methods("GET")
}

func (tc *TodoController)  GetTodoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	name := mux.Vars(r)["Name"]
	todo:= &models.Todo{}
	//var todo models.Todo
	err := tc.db.Find(bson.M{"Name": name}, todo)
	log.Println("RES=>" + todo.Name)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{message: %q}", "Database error")
		log.Println("Failed getting task: ", err)
		return
	}
	if todo.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{message: %q}", "Task not found")
		return
	}
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//var todos []models.Todo
	todos := make([]models.Todo, 0, 10)
	err := tc.db.FindAll(todos)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{message: %q}", "Database error")
		log.Println("Failed getting tasks: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{message: %q}", "Incorrect body")
		log.Println("Incorrect Body from request : ", err)
		return
	}

	err = tc.db.Save(todo)
	if err != nil {
		fmt.Fprintf(w, "{message: %q}", "Database error")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed insert task: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path + "/" + todo.Name)
	w.WriteHeader(http.StatusCreated)
}