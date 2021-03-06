package question

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

type Question struct {
	Id         string     `json:"id"`
	Text       string     `json:"text"`
	Tag        string     `json:"tag"`
	Order      string     `json:"order"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Id      string              `json:"id"`
	Text    string              `json:"text"`
	Order   string              `json:"order"`
	Options []map[string]string `json:"options"`
}

func Init(router *mux.Router) {
	router.Methods("GET").Path("/questions").HandlerFunc(listQuestions)
	router.Methods("GET").Path("/questions/{id}").HandlerFunc(listQuestionsId)
}

func listQuestions(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)

	rows, err := db.Query("SELECT id, text, tag, display_order FROM questions")
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		defer rows.Close()

		var id, text, tag, order string
		questions := make([]map[string]string, 0)
		for rows.Next() {
			rows.Scan(&id, &text, &tag, &order)
			questions = append(questions, map[string]string{"id": id, "text": text, "tag": tag, "order": order})
		}

		render.JSON(w, http.StatusOK, questions)
	}
}

func listQuestionsId(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)
	id := mux.Vars(req)["id"]

	question, err := getQuestionById(db, id)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rows, err := db.Query("SELECT c.id, c.text, c.display_order, o.id, o.text, o.display_order FROM categories c LEFT JOIN options o ON c.id = o.category_id WHERE c.question_id = $1 ORDER BY c.display_order, o.display_order", id)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		defer rows.Close()

		var cid, ctext, corder, pid string
		var oid, otext, oorder sql.NullString
		categories := make([]Category, 0)
		options := make([]map[string]string, 0)

		for rows.Next() {
			rows.Scan(&cid, &ctext, &corder, &oid, &otext, &oorder)

			if cid != pid {
				options = make([]map[string]string, 0)
				pid = cid
			} else {
				categories = categories[:len(categories)-1]
			}

			if oid.Valid {
				options = append(options, map[string]string{"id": oid.String, "text": otext.String, "order": oorder.String})
			}

			categories = append(categories, Category{cid, ctext, corder, options})
		}

		question.Categories = categories
		render.JSON(w, http.StatusOK, question)
	}
}

func getQuestionById(db *sql.DB, id string) (Question, error) {
	var question Question
	rows, err := db.Query("SELECT id, text, tag, display_order FROM questions WHERE id = $1", id)
	if err != nil {
		return question, err
	} else {
		var id, text, tag, order string
		if rows.Next() {
			rows.Scan(&id, &text, &tag, &order)
			question = Question{id, text, tag, order, nil}
		}
	}

	return question, nil
}
