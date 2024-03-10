package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"web_go/src/app"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)

}
