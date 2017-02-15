package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/deoliveiraromain/todo_api/db"
	"github.com/deoliveiraromain/todo_api/models"
	"github.com/deoliveiraromain/todo_api/repositories"
	"github.com/gorilla/mux"
)

type TodoController struct {
	db db.Mongo
}

func NewTodoController(db db.Mongo) *TodoController {
	tc := &TodoController{db: db}
	return tc
}
func (tc *TodoController) Register(router *mux.Router) {
	router.HandleFunc("/tasks", tc.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{name}", tc.GetTaskByName).Methods("GET")
	router.HandleFunc("/task", tc.CreateTodo).Methods("POST")
}

func (tc *TodoController) GetTaskByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.TaskRepo{session.DB(tc.db.DatabaseName).C("todo")}

	name := mux.Vars(r)["name"]
	log.Println("Search Task By Name =>" + name)
	task, err := repo.FindByName(name)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "{message: %q}", "Task not found")
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{message: %q}", "Database error")
			log.Println("Failed getting task: ", err)
			return
		}
	}
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *TodoController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.TaskRepo{session.DB(tc.db.DatabaseName).C("todo")}
	tasks, err := repo.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{message: %q}", "Database error")
		log.Println("Failed getting tasks: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var task models.TaskResource
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{message: %q}", "Incorrect body")
		log.Println("Incorrect Body from request : ", err)
		return
	}

	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.TaskRepo{session.DB(tc.db.DatabaseName).C("todo")}

	err = repo.Create(&task.Data)
	if err != nil {
		fmt.Fprintf(w, "{message: %q}", "Database error")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed insert task: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+task.Data.Name)
	json.NewEncoder(w).Encode(task)
	w.WriteHeader(http.StatusCreated)
}
