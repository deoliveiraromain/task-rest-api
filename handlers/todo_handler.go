package handlers

import (
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/deoliveiraromain/todo_api/models"
)

type TodoController struct {
	db db.DB
}

func NewTodoController(db db.DB) *TodoController {
	tc := &TodoController{db :db}
	return tc
}
func (tc *TodoController) Register(router *mux.Router) {
	router.HandleFunc("/todos", tc.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", tc.GetTodo).Methods("GET")
}

func (tc *TodoController)  GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(models.Todo{
		1,
		"Write presentation",
		false,
		//"0001-01-01T00:00:00Z",
	})
	if err != nil {
		panic(err)
	}
}

func (tc *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(models.Todo{
		1,
		"Write presentation",
		false,
		//"0001-01-01T00:00:00Z",
	})
	if err != nil {
		panic(err)
	}
}