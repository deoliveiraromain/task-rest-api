package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deoliveiraromain/task-rest-api/db"
	"github.com/deoliveiraromain/task-rest-api/models"
	"github.com/deoliveiraromain/task-rest-api/repositories"
	"github.com/gorilla/mux"
)

//TodoController : struct for task Controller with mongo repo
type TaskController struct {
	db db.Mongo
}

func NewTaskController(db db.Mongo) *TaskController {
	tc := &TaskController{db: db}
	return tc
}
func (tc *TaskController) Register(router *mux.Router) {
	router.HandleFunc("/tasks", tc.getAllTasks).Methods("GET")
	router.HandleFunc("/task/{name}", tc.getTaskByName).Methods("GET")
	router.HandleFunc("/task", tc.createTask).Methods("POST")
	router.HandleFunc("/task/{name}", tc.updateTask).Methods("POST")
}

func (tc *TaskController) getTaskByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.NewTaskRepo(session, tc.db.DatabaseName)

	name := mux.Vars(r)["name"]
	log.Println("search Task By Name =>" + name)
	task, err := repo.FindByName(name)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			log.Println("Task "+name+"not found", err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed getting task: ", err)
		return
	}
	if err = json.NewEncoder(w).Encode(task); err != nil {
		panic(err)
	}
}

func (tc *TaskController) getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.NewTaskRepo(session, tc.db.DatabaseName)
	tasks, err := repo.All()
	if err != nil {
		log.Println("Failed getting tasks: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}

func (tc *TaskController) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.TaskResource
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Incorrect Body from request : ", err)
		return
	}

	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.NewTaskRepo(session, tc.db.DatabaseName)

	if err = repo.Create(&task.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed insert task: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Location", r.URL.Path+"/"+task.Data.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func (tc *TaskController) updateTask(w http.ResponseWriter, r *http.Request) {
	var task models.TaskResource
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Incorrect Body from request : ", err)
		return
	}
	session := tc.db.Session.Copy()
	defer session.Close()
	repo := repositories.NewTaskRepo(session, tc.db.DatabaseName)

	name := mux.Vars(r)["name"]
	log.Println("search Task By Name =>" + name)
	taskDb, err := repo.FindByName(name)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			log.Println("Task "+name+"not found", err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed getting task: ", err)
		return
	}

	if err = repo.Update(&taskDb.Data, &task.Data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed update task: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Location", r.URL.Path+"/"+task.Data.Name)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}
