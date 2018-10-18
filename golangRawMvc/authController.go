package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

//ROUTES
const (
	register = "/api/auth/register"
	about    = "/about"
	contact  = "/contact"
)

type AuthController struct {
	kek string
}

func (auth *AuthController) initAuthController(mux *httprouter.Router) {
	mux.POST(register, auth.Register)
}

func (auth *AuthController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("1")
	var username, password string
	if username, password = r.FormValue("username"), r.FormValue("password"); username == "" && password == "" {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}
	fmt.Println("2")

	hashedSalt := hashAndSalt(password)
	userID := createUser(hashedSalt, username)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello, you've created new user: %v", userID)

}
func hashAndSalt(pwd string) []byte {

	bytesPwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytesPwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}
func verifyPassword(storedPwd []byte, incomingPwd string) bool {

	byteIncomingPwd := []byte(incomingPwd)
	err := bcrypt.CompareHashAndPassword(storedPwd, byteIncomingPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func createUser(pwd []byte, username string) int64 {
	db, err := sql.Open("mysql", "newuser:password@/localhost:3306")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO Users (Username, PasswordHash) VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(username, pwd)
	if err != nil {
		panic(err.Error())
	}
	userID, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("%v", err.Error)
	}
	return userID
}
