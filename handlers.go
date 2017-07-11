package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

/*
checkError is a function
*/
func checkError(err error) {
	fmt.Println(err.Error())
	fmt.Println("------------------------------------------")
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

/*
TokenRequest is a struct
*/
type TokenRequest struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	UserID string  `json:"userid"`
}

/*
GetToken is a very important function, it makes sure the user is within
our geofences and not on some sort of school campus, it also assigns roles
to people defining what they can and cannot do
*/
func GetToken(w http.ResponseWriter, req *http.Request) {
	// First we should parse to see if they are within
	// the bounds of the geofence or even gave us a lat
	// and long
	var tr TokenRequest
	json.NewDecoder(req.Body).Decode(&tr)
	fmt.Println(tr.Lat)
	fmt.Println(tr.Lon)
	// Make sure they gave us a lat and long
	if tr.Lat == 0 || tr.Lon == 0 {
		json.NewEncoder(w).Encode(Response{"Not within geofence", "invalid_location"})
		return
	}
	var point string
	point = "'POINT(" + strconv.FormatFloat(tr.Lat, 'f', 6, 64) + " " + strconv.FormatFloat(tr.Lon, 'f', 6, 64) + ")'"

	rows, err := db.Query("SELECT ST_CONTAINS(fence, PointFromText(?)) as valid FROM geofences order by valid desc limit 1", point)

	if err != nil {
		// Internal server error
		fmt.Println(err.Error())
		return
	}
	var truefalse int
	// Grab the value we want
	for rows.Next() {
		rows.Scan(&truefalse)
	}

	if truefalse == 1 {
		// They are within atleast one geofence, yay!
		// Now we have to check for their userid
	} else {
		// They are not within it so return a message telling
		// them to go to a real school
		json.NewEncoder(w).Encode(Response{"Not within geofence", "invalid_location"})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["userid"] = "DJ093L98JS"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(signingKey)

	// Create our authorization cookie with the new token
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().AddDate(0, 0, 1),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	var resp Response
	resp.Message = "success"
	json.NewEncoder(w).Encode(resp)
}

/*
GetAllMessages is a function
*/
func GetAllMessages(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT messages.id, messages.message, " +
		"coalesce(sum(votes.updoot) - sum(votes.downdoot),0) " +
		"AS updoots FROM messages LEFT JOIN votes on messages.id = votes.message " +
		"group by messages.id order by updoots desc")
	// If we experience some kind of error
	if err != nil {
		checkError(err)
		w.WriteHeader(500)
		w.Write([]byte("Uh oh!"))
		return
	}

	var id int
	var mess string
	var ups int
	// These are the messages we will
	// be sending back
	var messages []Message
	//fmt.Println(rows)
	for rows.Next() {
		var message Message
		rows.Scan(&id, &mess, &ups)
		message.ID = id
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

	if message.Message == "" {
		var empty Empty
		json.NewEncoder(w).Encode(empty)
	} else {
		json.NewEncoder(w).Encode(message)
	}
}

/*
CreateMessage is a function
*/
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
