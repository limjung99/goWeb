package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	// using anotation for json !!!
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	user := new(User)
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreateAt = time.Now()
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &fooHandler{})

	mux.HandleFunc("/handleFunc", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, this is handleFunc!")
	})

	http.ListenAndServe(":8080", mux)
}
