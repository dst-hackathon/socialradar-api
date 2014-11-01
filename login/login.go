package login

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"github.com/dgrijalva/jwt-go"
	"github.com/dst-hackathon/socialradar-api/configuration"
	"net/http"
	"encoding/json"
)

type loginInfo struct {
	Email string
	Id string
	Password string
}

var signKey interface{} = []byte{1,1,1,1,1,1,1}
	
func Init(router *mux.Router) {
	router.Methods("POST").Path("/login").HandlerFunc(loginHandler)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)
	config := context.Get(req, "config").(configuration.Configuration)
	
	// Keep Sign Key
	signKey = []byte(config.ApiSignKey)
	
    inputInfo := parseRequestBody(req)
	if inputInfo.Email == "" || inputInfo.Password == "" {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Please enter email and password"})
		return;
	}
	userInfo, err := getUser(db, inputInfo.Email)

	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return;
	}

	response := validateUser(inputInfo, userInfo)
	render.JSON(w, http.StatusOK, response)
}

func parseRequestBody(req *http.Request) loginInfo {
    var info loginInfo
	decoder := json.NewDecoder(req.Body)
    decoder.Decode(&info)
	return info
}

func getUser(db *sql.DB, email string) (loginInfo, error) {
	var userInfo loginInfo

	tx, err := db.Begin()
	if err != nil {
		return userInfo, err
	}

	rows, err := tx.Query("SELECT email, id, encrypted_password FROM users where email = $1", email)
	if err != nil {
		tx.Rollback()
		return userInfo, err
	}

	defer rows.Close()
		
	var res_email, id, encrypted_password string
	rows.Next()
	rows.Scan(&res_email, &id, &encrypted_password)

	userInfo.Email = res_email
	userInfo.Id = id
	userInfo.Password = encrypted_password

	return userInfo, nil
}

func validateUser(inputInfo loginInfo, actualInfo loginInfo) map[string]string {
	if (inputInfo.Password != actualInfo.Password) {
		return map[string]string{"error":"Invalid username or password"}
	}
	token, err := createToken(actualInfo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return map[string]string{"success":"success", "token":token, "id":actualInfo.Id}
}

func createToken(actualInfo loginInfo) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    token.Claims["id"] = actualInfo.Id
	return token.SignedString(signKey)
}

func decodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return signKey, nil
    })
	if err == nil && token.Valid {
		return token, nil
    }
	return nil, err
}
