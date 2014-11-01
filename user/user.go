package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dst-hackathon/socialradar-api/configuration"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"strconv"
	"sort"
)

func Init(router *mux.Router) {
	router.Methods("POST").Path("/users/{id}/answer").HandlerFunc(saveUserAnswer)
	router.Methods("GET").Path("/users/{id}/answer").HandlerFunc(getUserAnswer)
	router.Methods("POST").Path("/users/{id}/avatar").HandlerFunc(postAvatar)
	router.Methods("GET").Path("/users/{id}/avatar").HandlerFunc(getAvatar)
	router.Methods("GET").Path("/users/{id}/friendsuggestions").HandlerFunc(suggestFriends)
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

func getUserAnswer(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)
	userId := mux.Vars(req)["id"]

	rows, err := db.Query(`
		SELECT c.question_id, uc.category_id, uo.option_id
		FROM users_categories uc
		JOIN categories c ON uc.category_id = c.id
		LEFT OUTER JOIN users_options uo ON uo.user_id = $1 AND uo.option_id IN (
				select o.id FROM options o
				WHERE o.category_id = uc.category_id
			)
		WHERE uc.user_id = $1
		ORDER BY c.question_id, uc.category_id, uo.option_id
		`, userId)

	if err != nil {
		log.Fatal(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		defer rows.Close()

		result := make(map[string]map[string][]int)

		var qidNS, cidNS, oidNS sql.NullString
		var qid, cid, oid string
		for rows.Next() {
			rows.Scan(&qidNS, &cidNS, &oidNS)

			if qidNS.Valid {
				qid = qidNS.String
			}

			if cidNS.Valid {
				cid = cidNS.String
			}

			if oidNS.Valid {
				oid = oidNS.String
			} else {
				oid = ""
			}

			if result[qid] == nil {
				result[qid] = make(map[string][]int)
			}

			if result[qid][cid] == nil {
				result[qid][cid] = make([]int, 0)
			}

			if oid != "" {
				oidInt, _ := strconv.Atoi(oid)
				result[qid][cid] = append(result[qid][cid], oidInt)
			}
		}

		render.JSON(w, http.StatusOK, result)
	}
}

type PostAvatarResult struct {
	Status   string
	Filename string
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
	fileExtension := filenameSplits[len(filenameSplits)-1]
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

func getAvatar(w http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
	config := context.Get(req, "config").(configuration.Configuration)
	db := context.Get(req, "db").(*sql.DB)
	render := context.Get(req, "render").(*render.Render)

	rows, err := db.Query("SELECT avatar_path FROM users WHERE id = $1", userId)
	if err != nil {
		render.Data(w, http.StatusBadRequest, nil)
	} else {
		defer rows.Close()

		var avatar_path sql.NullString

		for rows.Next() {
			rows.Scan(&avatar_path)
		}

		if avatar_path.Valid {
			filepath := config.AvatarPath + avatar_path.String
			http.ServeFile(w, req, filepath)
		} else {
			render.Data(w, http.StatusBadRequest, nil)
		}
	}
}

func suggestFriends(w http.ResponseWriter, req *http.Request) {
	db := context.Get(req, "db").(*sql.DB)
	render := context.Get(req, "render").(*render.Render)
	var user_id string = mux.Vars(req)["id"]
	
	resultByOptions, err := calculateByUserOptions(db, user_id)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resultByCategories, err := calculateByUserCategories(db, user_id)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	mergeResult := groupResult(resultByOptions, resultByCategories)
	response := make([]map[string]string, 0)
	for _, f := range mergeResult {
		response = append(response, map[string]string{"id":f.id, "weight":strconv.Itoa(f.weight), "email":f.email})
	}
	
	render.JSON(w, http.StatusOK, response)
}

type frientList []friend
type friend struct {
	id string
	weight int
	email string
}
func (s frientList) Len() int { return len(s) }
func (s frientList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s frientList) Less(i, j int) bool { 
	return s[i].weight < s[j].weight
}

func groupResult(resultByOptions []friend, resultByCategories []friend) []friend {
	result := map[string]friend{}
	for _, f := range resultByOptions {
		result[f.id] = f
	}
	for _, f := range resultByCategories {
		if result[f.id].id != f.id {
			result[f.id] = f
		} else {
			oldF := result[f.id]
			oldF.weight += f.weight
			result[f.id] = oldF
		}
	}

	sortResult := make([]friend, 0)
	for _, f := range result {
		sortResult = append(sortResult, f)
	}
	sort.Sort(sort.Reverse(frientList(sortResult)))
	return sortResult
}

func calculateByUserCategories(db *sql.DB, user_id string) ([]friend, error) {
	rows, err := db.Query(
		`SELECT friend.id, count(fuo) as weight, friend.email 
		FROM USERS friend 
		INNER JOIN USERS_CATEGORIES fuo ON fuo.user_id = friend.id 
		WHERE friend.id <> $1 
		AND EXISTS 
		(select 'Y' from USERS_CATEGORIES cuo 
		WHERE cuo.category_id = fuo.category_id 
		AND cuo.user_id = $2) 
		GROUP BY friend.id 
		ORDER BY weight DESC;`, user_id, user_id)

	if err != nil {
		return []friend{}, err
	}
	defer rows.Close()

	var id, weight, email string
	suggestions := make([]friend, 0)
	for rows.Next() {
		rows.Scan(&id, &weight, &email)
		iWeight, _ := strconv.Atoi(weight)
		suggestions = append(suggestions, friend{id: id, weight: iWeight, email: email})
	}
	return suggestions, nil
}

func calculateByUserOptions(db *sql.DB, user_id string) ([]friend, error) {
	rows, err := db.Query(
		`SELECT friend.id, count(fuo) as weight, friend.email 
		FROM USERS friend 
		INNER JOIN USERS_OPTIONS fuo ON fuo.user_id = friend.id 
		WHERE friend.id <> $1 
		AND EXISTS 
		(select 'Y' from USERS_OPTIONS cuo 
		WHERE cuo.option_id = fuo.option_id 
		AND cuo.user_id = $2) 
		GROUP BY friend.id 
		ORDER BY weight DESC;`, user_id, user_id)

	if err != nil {
		return []friend{}, err
	}
	defer rows.Close()

	var id, weight, email string
	suggestions := make([]friend, 0)
	for rows.Next() {
		rows.Scan(&id, &weight, &email)
		iWeight, _ := strconv.Atoi(weight)
		suggestions = append(suggestions, friend{id: id, weight: iWeight, email: email})
	}
	return suggestions, nil
}
