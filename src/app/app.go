package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func indexHadler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTestTodos() {
	todoMap[1] = &Todo{
		1, "Buy a milk", true, time.Now(),
	}
	todoMap[2] = &Todo{
		2, "Learn Go", false, time.Now(),
	}
	todoMap[3] = &Todo{
		2, "study", false, time.Now(),
	}
}

func addTodoListHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	rd.JSON(w, http.StatusOK, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestTodos()
	rd = render.New()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHadler)
	router.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	router.HandleFunc("/todos", addTodoListHandler).Methods("POST")
	router.HandleFunc("todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	router.HandleFunc("complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	return router
}
