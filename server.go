package main

import (
	"database/sql"
	"github.com/codegangsta/negroni"
	"github.com/dst-hackathon/socialradar-api/question"
	"github.com/dst-hackathon/socialradar-api/user"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	question.Init(router)
	user.Init(router)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(RenderIntializer))
	n.Use(negroni.HandlerFunc(DbInitializer))
	n.UseHandler(router)
	n.Run(":3000")
}

func RenderIntializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	render := render.New(render.Options{
		IndentJSON: true,
	})
	context.Set(r, "render", render)

	next(rw, r)
}

func DbInitializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/socialradar_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		http.Error(rw, err.Error(), 500)
		return
	}

	context.Set(r, "db", db)
	next(rw, r)

	db.Close()
}
