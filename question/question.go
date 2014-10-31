package question

import (
	"database/sql"
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
	db := context.Get(req, "db").(*sql.DB)

	rows, err := db.Query("SELECT * FROM questions")
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		defer rows.Close()

		var id, text, order string
		questions := make([]map[string]string, 0)
		for rows.Next() {
			rows.Scan(&id, &text, &order)
			questions = append(questions, map[string]string{"id": id, "text": text, "order": order})
		}

		render.JSON(w, http.StatusOK, questions)
	}

}
