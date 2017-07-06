package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

/*
checkError is a function
*/
func checkError(err error) {
	fmt.Println(err.Error())
	fmt.Println("------------------------------------------")
}

/*
GetAllMessages is a function
*/
func GetAllMessages(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM messages")
	// If we experience some kind of error
	if err != nil {
		checkError(err)
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
	if messages == nil {
		messages = make([]Message, 0)
	}
	// else we should have our rows
	json.NewEncoder(w).Encode(messages)
}

/*
GetMessage is a function
*/
func GetMessage(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	rows, err := db.Query("SELECT * FROM messages where id=?", params["id"])
	if err != nil {
		checkError(err)
	}
	var message Message
	var id int
	var mess string
	var userid string
	var ups int
	for rows.Next() {
		rows.Scan(&id, &mess, &userid, &ups)

		message.Message = mess
		message.Updoots = ups
	}
	if err != nil {
		checkError(err)
	}

	if(message.Message == "") {
		var empty Empty
		json.NewEncoder(w).Encode(empty)
	} else {
		json.NewEncoder(w).Encode(message)
	}
}
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
