package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/deoliveiraromain/todo_api/db"
	"log"
)

type TodoController struct {
	db db.DB
	//routes []routes.Route
}

func (tc *TodoController) GetRoutes() Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			tc.Index,
		},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			tc.TodoIndex,
		},
		Route{
			"TodoShow",
			"GET",
			"/todo/{todoId}",
			tc.TodoShow,
		},
	}
}

func NewTodoController(db db.DB) *TodoController {
	tc := &TodoController{db :db}
	return tc
}

func (tc *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	log.Print("COUCOU")
	fmt.Fprint(w, "Welcome!\n")
}

func (tc *TodoController)  TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("") // TODO : a remplacer
	if err != nil {
		panic(err)
	}
}

func (tc *TodoController) TodoShow(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//var todoId int
	//var err error
	//if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
	//	panic(err)
	//}
	//todo := RepoFindTodo(todoId)
	//if todo.Id > 0 {
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.StatusOK)
	//	if err := json.NewEncoder(w).Encode(todo); err != nil {
	//		panic(err)
	//	}
	//	return
	//}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
	//	panic(err)
	//}

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
//func TodoCreate(w http.ResponseWriter, r *http.Request) {
//	var todo Todo
//	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//	if err != nil {
//		panic(err)
//	}
//	if err := r.Body.Close(); err != nil {
//		panic(err)
//	}
//	if err := json.Unmarshal(body, &todo); err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(422) // unprocessable entity
//		if err := json.NewEncoder(w).Encode(err); err != nil {
//			panic(err)
//		}
//	}
//
//	t := RepoCreateTodo(todo)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusCreated)
//	if err := json.NewEncoder(w).Encode(t); err != nil {
//		panic(err)
//	}
//}