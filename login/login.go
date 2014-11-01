package login

import (
	"database/sql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	//"log"
)

type loginInfo struct {
	Id string
	Password string
	Uid string
}

var signKey interface{} = []byte{1,1,1,1,1,1,1}
	
func Init(router *mux.Router) {
	router.Methods("POST").Path("/login").HandlerFunc(loginHandler)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)

    inputInfo := parseRequestBody(req)
	userInfo, err := getUser(db, inputInfo.Uid)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return;
	}

	//log.Printf("input uid: " + inputInfo.Uid)
	//log.Printf("input pass: " + inputInfo.Password + ", get pass: " + userInfo.Password)

	response := validateUser(inputInfo, userInfo)
	render.JSON(w, http.StatusOK, response)
	
	/*token, err := decodeToken(response["token"])
	if err == nil {
		log.Printf("decode token.Claims[id] = " + token.Claims["id"].(string))
	} else {
		log.Print("Invalid token")
	}*/
}

func parseRequestBody(req *http.Request) loginInfo {
    var info loginInfo
	decoder := json.NewDecoder(req.Body)
    decoder.Decode(&info)
	return info
}

func getUser(db *sql.DB, userId string) (loginInfo, error) {
	var userInfo loginInfo

	tx, err := db.Begin()
	if err != nil {
		return userInfo, err
	}

	rows, err := tx.Query("SELECT id, encrypted_password, uid FROM users where uid = $1", userId)
	if err != nil {
		tx.Rollback()
		return userInfo, err
	}

	defer rows.Close()
		
	var id, encrypted_password, uid string
	rows.Next()
	rows.Scan(&id, &encrypted_password, &uid)

	userInfo.Id = id
	userInfo.Password = encrypted_password
	userInfo.Uid = uid

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
	return map[string]string{"success":"success", "token":token}
}

func createToken(actualInfo loginInfo) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    token.Claims["id"] = actualInfo.Id

	//log.Printf("token.Claims[id] = " + token.Claims["id"].(string))
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

func addUserHandler(w http.ResponseWriter, req *http.Request) {
	render := context.Get(req, "render").(*render.Render)
	db := context.Get(req, "db").(*sql.DB)

	_, err := db.Query("INSERT INTO users(encrypted_password,email,uid) values ('mypass','pongsanti.tanvejsilp@gmail.com','myid')")
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else {
		render.JSON(w, http.StatusOK, map[string]string{"success": ""})
	}
}
