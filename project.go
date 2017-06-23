package main

import (
	"encoding/json"
	"net/http"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
Message is a struct
*/
type Message struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  string `json:"userid,omitempty"`
}

// Declare our global variables in
// place of our database
var messages []Message

/*
GetAllMessages is a function
*/
func GetAllMessages(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(messages)
}

/*
GetMessage is a function
*/
func GetMessage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range messages {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// If we don't find a person with that id
	// send back a blank object
	json.NewEncoder(w).Encode(&Message{})
}

/*
CreateMessage is a function
*/
func CreateMessage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// declare a new message
	var message Message
	_ = json.NewDecoder(req.Body).Decode(&message)
	message.ID = params["id"]
	messages = append(messages, message)
	// return the newly created object
	json.NewEncoder(w).Encode(message)
}

/*
DeleteMessage is a function
*/
func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range messages {
		// loop through until the id matches,
		// then remove it and break
		if item.ID == params["id"] {
			messages = append(messages[:index], messages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(messages)
}

/*
main is a function
*/
func main() {
	db, err := sql.Open("mysql")
	router := mux.NewRouter()
	messages = append(messages, Message{ID: "1", Message: "The very first post!"})

	router.HandleFunc("/message", GetAllMessages).Methods("GET")
	router.HandleFunc("/message/{id}", GetMessage).Methods("GET")
	router.HandleFunc("/message/{id}", CreateMessage).Methods("POST")
	router.HandleFunc("/message/{id}", DeleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
