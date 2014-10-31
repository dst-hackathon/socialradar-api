package user

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

func Init(router *mux.Router) {
	router.Methods("POST").Path("/users/{id}/answer").HandlerFunc(saveUserAnswer)
}

/*
Example request is:

{
  "1": {  // Question ID
    "1": [1, 2],  // "Selected Category": [Selected Option ID, ....]
    "2": []
  },
  "2": {
    "5": [10]
  }
}
*/
func saveUserAnswer(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)
	decoder := json.NewDecoder(req.Body)
	userId := mux.Vars(req)["id"]

	var data map[string]map[string][]int
	err := decoder.Decode(&data)
	if err == nil {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		for _, qvalue := range data {
			for cid, cvalue := range qvalue {
				_, err := tx.Exec("INSERT INTO users_categories(user_id, category_id) VALUES ($1, $2)", userId, cid)

				if err != nil {
					log.Fatal(err)
					tx.Rollback()
					render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
					return
				}

				for _, oid := range cvalue {
					_, err := tx.Exec("INSERT INTO users_options(user_id, option_id) VALUES ($1, $2)", userId, oid)

					if err != nil {
						log.Fatal(err)
						tx.Rollback()
						render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
						return
					}
				}
			}
		}

		render.JSON(w, http.StatusOK, map[string]bool{"success": true})
		tx.Commit()
	} else {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
}
