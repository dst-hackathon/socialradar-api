package question

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func Init(router *mux.Router) {
	router.Methods("GET").Path("/question").HandlerFunc(listQuestion)
}

func listQuestion(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)

	render.JSON(w, http.StatusOK, map[string]string{"hello": "Hello"})
}
