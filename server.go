package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dst-hackathon/socialradar-api/configuration"
	"github.com/dst-hackathon/socialradar-api/login"
	"github.com/dst-hackathon/socialradar-api/question"
	"github.com/dst-hackathon/socialradar-api/signup"
	"github.com/dst-hackathon/socialradar-api/user"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"strings"
)

var config configuration.Configuration = configuration.ReadFile()
var signKey interface{} = []byte{1, 1, 1, 1, 1, 1, 1}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	question.Init(router)
	user.Init(router)
	login.Init(router)
	signup.Init(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(RenderIntializer))
	n.Use(negroni.HandlerFunc(ConfigInitializer))
	n.Use(negroni.HandlerFunc(DbInitializer))
	n.Use(c)
	n.Use(negroni.HandlerFunc(SecurityInitializer))
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

func ConfigInitializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "config", config)
	next(rw, r)
}

func DbInitializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/socialradar_development?sslmode=disable", config.DbUser, config.DbPassword, config.DbHost, config.DbPort)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	context.Set(r, "db", db)
	next(rw, r)

	db.Close()
}

func SecurityInitializer(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var tokenString string
	uri := r.RequestURI
	log.Printf("Security checking for : " + uri)
	if strings.HasPrefix(uri, "/login") ||
		strings.HasPrefix(uri, "/signup") {
		next(rw, r)
		return
	}
	tokenString = r.Header.Get("token") //r.FormValue("token")
	log.Printf("Token provide:" + tokenString)
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			var signKey []byte = []byte(config.ApiSignKey)
			return signKey, nil
		})
		if err == nil && token.Valid {
			log.Printf("decode token.Claims[id] = " + token.Claims["id"].(string))
			context.Set(r, "user_id", token.Claims["id"])
		} else {
			log.Printf("Error")
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
	}
	next(rw, r)
}
