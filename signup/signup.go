package signup

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"encoding/json"
)

type Signup struct {
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
}

func Init(router *mux.Router) {
	router.Methods("POST").Path("/signup").HandlerFunc(signup)
}
/*
{
	"email": "test@email.com",
	"password": "xxfji"
}
*/
func signup(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)
	decoder := json.NewDecoder(req.Body)

	signup := Signup{} 
	err := decoder.Decode(&signup)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		rows, err := db.Query("SELECT 1 FROM users WHERE email = $1", signup.Email)
		if err != nil {
			render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			defer rows.Close()
			if rows.Next() {
				render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Email's already existed."})
				return
			}
		}

		stmt, err := db.Prepare("INSERT INTO users(encrypted_password, email) VALUES ($1, $2)")
		if err != nil {
			log.Print(err)
			render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		_, err = stmt.Exec(signup.Password, signup.Email)
		if err != nil {
			log.Print(err)
			render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		if err != nil {
			render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		} else {
			render.JSON(w, http.StatusOK, map[string]string{"success": ""})
		}
	}
}