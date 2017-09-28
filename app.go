package main

import (
	"net/http"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
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
	db, err = sql.Open("postgres", "postgres://nuser:npassword@tcp(updoots_db:5432)/testdb?charset=utf8mb4&parseTime=true")
	if err != nil {
		return
	}

	defer db.Close()

	router := APIRouter()
	// Make sure to allow all requests

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.66.50:8100","http://127.0.0.1:8100","https://app.updoot.us","https://updoot.us"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)
        
	log.Fatal(http.ListenAndServe(":3001", handler))
}
