package main

import (
	"encoding/json"
	"net/http"

	"database/sql"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

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

// Declare our global variables in
// place of our database
var db *sql.DB
var err error

/*
GetAllMessages is a function
*/
func GetAllMessages(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM messages")
	// If we experience some kind of error
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Uh oh!"))
		return
	}

	var id int
	var mess string
	var userid string
	var ups int
	// These are the messages we will
	// be sending back
	var messages []Message
	//fmt.Println(rows)
	for rows.Next() {
		var message Message
		rows.Scan(&id, &mess, &userid, &ups)

		message.Message = mess
		message.Updoots = ups
		messages = append(messages, message)
	}
	// Makes sure the client sees application/json
	w.WriteHeader(http.StatusOK)
	// else we should have our rows
	json.NewEncoder(w).Encode(messages)
}

/*
checkError is a function
*/
func checkError(err error) {
	fmt.Println(err.Error())
}

/*
GetMessage is a function
*/
func GetMessage(w http.ResponseWriter, req *http.Request) {
	return
	params := mux.Vars(req)
	stmt, err := db.Prepare("SELECT * FROM messages where id=?")
	checkError(err)
	res, err := stmt.Exec(params["id"])
	checkError(err)
	fmt.Println(res)
	// If we don't find a person with that id
	// send back a blank object
	json.NewEncoder(w).Encode(&Message{})
}

/*
CreateMessage is a function
*/
func CreateMessage(w http.ResponseWriter, req *http.Request) {
	return
	var nMessage Newmessage
	json.NewDecoder(req.Body).Decode(&nMessage)
	// declare a new message
	// All new messages will have 0 updoots to start
	stmt, err := db.Prepare("INSERT messages SET message=?,userid=?,updoots=0")
	if err != nil {
		return
	}
	res, err := stmt.Exec(nMessage.Message, "Y87YUHG989839RW09U98")
	checkError(err)
	// return the newly created object
	json.NewEncoder(w).Encode(res)
}

/*
DeleteMessage is a function
*/
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

/*
CreateTable is a function
*/
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
}

/*
main is a function
*/
func main() {
	// Note that here we must use a strict = rather than :=
	db, err = sql.Open("mysql", "nuser:npassword@tcp(pittyak_db:3306)/testdb")

	//checkError(err)
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/message", GetAllMessages).Methods("GET")
	router.HandleFunc("/message/{id}", GetMessage).Methods("GET")
	router.HandleFunc("/message", CreateMessage).Methods("POST")
	router.HandleFunc("/message/{id}", DeleteMessage).Methods("DELETE")
	router.HandleFunc("/admin/create", CreateTable).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", router))
}
