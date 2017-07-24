package main

import (
	"net/http"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

/*
signingKey is our global sercret shhhhhhh
*/
var signingKey = []byte("super duper super secure lollipop")

// Declare our global variables in
// place of our database
var db *sql.DB
var err error

/*
main is a function
*/
func main() {
	fmt.Println("+----------------------+")
	fmt.Println("| SERVER HAS RESTARTED |")
	fmt.Println("+----------------------+")

	// Note that here we must use a strict = rather than :=
	db, err = sql.Open("mysql", "nuser:npassword@tcp(updoots_db:3306)/testdb")
	if err != nil {
		return
	}

	defer db.Close()

	router := APIRouter()
	// Make sure to allow all requests

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH", "DELETE"})
        
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
