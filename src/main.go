package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web_go/src/decoHandler"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("{LOGGER1} Started")
	h.ServeHTTP(w, r)
	log.Println("{LOGGER1} Completed", time.Since(start).Milliseconds())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	newMux := decoHandler.NewDecoHandler(mux, logger)
	return newMux
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
