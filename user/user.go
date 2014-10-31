package user

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/dst-hackathon/socialradar-api/configuration"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
)

func Init(router *mux.Router) {
	router.Methods("POST").Path("/users/{id}/answer").HandlerFunc(saveUserAnswer)
	router.Methods("POST").Path("/users/{id}/avatar").HandlerFunc(postAvatar)
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

type PostAvatarResult struct {
	Status 		string
	Filename 	string
}

func postAvatar(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	userId := mux.Vars(req)["id"]

	file, header, err := req.FormFile("file")
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	originalFilename := header.Filename
	filenameSplits := strings.Split(originalFilename, ".")
	fileExtension := filenameSplits[len(filenameSplits) - 1]
	savedFilename := userId + "." + fileExtension

	config := context.Get(req, "config").(configuration.Configuration)
	out, err := os.Create(config.AvatarPath + savedFilename)
	if err != nil {
		fmt.Fprintln(w, "Unalble to create file.")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	// Update users.avatar_path
	db := context.Get(req, "db").(*sql.DB)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET avatar_path = $1 WHERE id = $2")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(savedFilename, userId)
	if err != nil {
		log.Fatal(err)
	}

	result := PostAvatarResult{Status: "Success", Filename: originalFilename}
	render.JSON(w, http.StatusOK, result)
}
