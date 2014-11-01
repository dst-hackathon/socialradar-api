package signup

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/dst-hackathon/socialradar-api/configuration"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"strings"
	"os"
	"io"
)

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

	file, _, err := req.FormFile("file")
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Image file is required."})
		return
	}
	defer file.Close()

	email := req.FormValue("email")
	password := req.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Email/Password cannot be null."})
		return
	}

	rows, err := db.Query("SELECT 1 FROM users WHERE email = $1", email)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
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
	_, err = stmt.Exec(password, email)
	if err != nil {
		log.Print(err)
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		err = postAvatar(req)
		render.JSON(w, http.StatusOK, map[string]string{"success": ""})
	}
	
}

func postAvatar(req *http.Request) (error) {
	db := context.Get(req, "db").(*sql.DB)
	email := req.FormValue("email")
	file, header, err := req.FormFile("file")
	defer file.Close()

	if err != nil {
		return err
	}

	var userId string
	rows, err := db.Query("SELECT id FROM users WHERE email = $1", email)

	if err != nil {
		return err
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&userId)
		}
	}

	originalFilename := header.Filename
	filenameSplits := strings.Split(originalFilename, ".")
	fileExtension := filenameSplits[len(filenameSplits) - 1]
	savedFilename := userId + "." + fileExtension

	config := context.Get(req, "config").(configuration.Configuration)
	out, err := os.Create(config.AvatarPath + savedFilename)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		//fmt.Fprintln(w, err)
		return err
	}

	// Update users.avatar_path
	if err != nil {
		log.Print(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET avatar_path = $1 WHERE id = $2")
	if err != nil {
		log.Print(err)
		return err
	}
	_, err = stmt.Exec(savedFilename, userId)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}