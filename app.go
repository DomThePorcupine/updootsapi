package main

import (
	"net/http"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
signingKey is our global sercret shhhhhhh
*/
var signingKey = []byte("super duper super secure lollipop")

/*
Claims is a struct
*/
type Claims struct {
	Expires int64  `json:"exp"`
	Admin   bool   `json:"admin"`
	UserID  string `json:"userid"`
}

/*
Message is a struct
*/
type Message struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  string `json:"userid,omitempty"`
	Updoots int    `json:"updoots"`
}

/*
Newmessage is a struct
*/
type Newmessage struct {
	Message string `json:"message,omitempty"`
	UserID  string `json:"userid,omitempty"`
}

/*
Database is a struct
*/
type Database struct {
	Key string `json:"key,omitempty"`
}

/*
Empty is a struct
*/
type Empty struct {
}

// Declare our global variables in
// place of our database
var db *sql.DB
var err error

/*

CreateMessage is a function

func CreateMessage(w http.ResponseWriter, req *http.Request) {
	var nMessage Newmessage
	json.NewDecoder(req.Body).Decode(&nMessage)
	// declare a new message
	// All new messages will have 0 updoots to start
	stmt, err := db.Prepare("INSERT messages SET message=?,userid=?,updoots=0")
	if err != nil {
		return
	}
	fmt.Println(nMessage.Message)
	res, err := stmt.Exec(nMessage.Message, "Y87YUHG989839RW09U98")
	if err != nil {
		checkError(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		checkError(err)
	}
	// return the newly created object
	json.NewEncoder(w).Encode(id)
}

DeleteMessage is a function
func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	return
	params := mux.Vars(req)

	stmt, err := db.Prepare("DELETE FROM messages WHERE id=?")
	checkError(err)
	res, err := stmt.Exec(params["id"])
	fmt.Println(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Uh oh"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Success."))
}


CreateTable is a function

func CreateTable(w http.ResponseWriter, req *http.Request) {
	return
	var keys Database
	json.NewDecoder(req.Body).Decode(&keys)
	// This is the key we will use to verify that they have acces
	// to the database, no point in implementing logins
	if keys.Key != "WKHEAW33X9ZJ9VAG5VTD" {
		w.WriteHeader(404)
		w.Write([]byte("404 page not found"))
		return
	}
}*/

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
		checkError(err)
	}

	defer db.Close()

	router := APIRouter()

	log.Fatal(http.ListenAndServe(":3001", router))
}
