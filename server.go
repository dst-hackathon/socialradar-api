package main

import (
	"github.com/codegangsta/negroni"
	"github.com/dst-hackathon/socialradar-api/question"
	"github.com/dst-hackathon/socialradar-api/user"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	question.Init(router)
	user.Init(router)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(RenderIntializer))
	n.UseHandler(router)
	n.Run(":3000")
}

func RenderIntializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	render := render.New(render.Options{})
	context.Set(r, "render", render)

	next(rw, r)
}
