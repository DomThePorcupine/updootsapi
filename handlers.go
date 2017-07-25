package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*
checkError is a function
*/
func checkError(err error) {
	fmt.Println(err.Error())
	fmt.Println("------------------------------------------")
}

type Claims struct {
    Expires int64  `json:"exp"`
	Admin   bool   `json:"admin"`
	UserID  string `json:"userid"`
    // recommended having
    jwt.StandardClaims
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
	
	if tr.UserID == "" {
		json.NewEncoder(w).Encode(Response{"No user id given", "invalid_id"})
		return
	}
	
	rows, err := db.Query("SELECT admin from users where userid = ?", tr.UserID)

	if err != nil {
		// Internal server error
		fmt.Println(err.Error())
		return
	}

	var truefalse int
	truefalse = -1
	// Grab the value we want
	for rows.Next() {
		rows.Scan(&truefalse)
	}

	// Declare the token we will be giving them
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	if truefalse == 1 {
		// It's brian or dom let him do whatever
		claims["admin"] = true
	} else if truefalse == 0 {
		// Regular user so just make sure they are not an admin
		claims["admin"] = false
	} else {
		json.NewEncoder(w).Encode(Response{"No user id given", "invalid_id"})
		return
	}

	// Set their userid
	claims["userid"] = tr.UserID
	// Make sure the token experies in a reasonable amount of time
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
	json.NewEncoder(w).Encode(Response{tokenString, "auth_successful"})
}

/*
GetAllMessages is a function
*/
func GetAllMessages(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("select id, messages.message, ifnull(doots,0) as totalvotes " + 
							"from messages left join(select votes.message, " +
							"cast((sum(votes.updoot) - sum(votes.downdoot)) as signed) " + 
							"as doots from votes group by votes.message) as votes " + 
							"on messages.id = votes.message having totalvotes > -3 order by ifnull(doots,0) desc")
	
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
	ups = -17
	// These are the messages we will
	// be sending back
	var messages []Message
	//fmt.Println(rows)
	for rows.Next() {
		var message Message
		rows.Scan(&id, &mess, &ups)
		if ups == -17 {
			message.Updoots = 0
		} else {
			message.Updoots = ups
		}
		message.ID = id
		message.Message = mess
		
		messages = append(messages, message)
		ups = -17
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
CreateMessage is a function
*/
func CreateMessage(w http.ResponseWriter, req *http.Request) {
	clms, ok := req.Context().Value(Claims{}).(Claims)
	if !ok {
		json.NewEncoder(w).Encode(Response{"invalid id", "invalid_id"})
		return
	}
	var nMessage Newmessage
	json.NewDecoder(req.Body).Decode(&nMessage)
	fmt.Println(nMessage.Message)
	// For now simply make sure we only keep 100 messages
	dl, err := db.Prepare(	"delete message from messages as message " + 
							"join(select created from messages order by created desc limit 1 offset 98)" + 
							" ctd on message.created < ctd.created;")
	if err != nil {
		return
	}
	dl.Exec()
	// declare a new message
	// All new messages will have 0 updoots to start
	stmt, err := db.Prepare("INSERT messages SET message=?,userid=?")
	if err != nil {
		return
	}
	
	res, err := stmt.Exec(nMessage.Message, clms.UserID)

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

/*
This is a function
*/

func DootOnMessage(w http.ResponseWriter, req *http.Request) {
	// Parse the post request
	var vote Vote
	json.NewDecoder(req.Body).Decode(&vote)
	fmt.Println("Message id")
	fmt.Println(vote.Message)
	// Check our claims
	clms, ok := req.Context().Value(Claims{}).(Claims)
	if !ok {
		json.NewEncoder(w).Encode(Response{"invalid id", "invalid_id"})
		return
	}
	if clms.Admin {
		// They can upvote as much as they want
		if vote.Doot == 1 {
			stmt, err := db.Prepare("INSERT votes SET message=?,userid=?,updoot=1")
			if err != nil {
				return
			}
			stmt.Exec(vote.Message, clms.UserID)
			return
		} else if vote.Doot == 0 {
			stmt, err := db.Prepare("INSERT votes SET message=?,userid=?,downdoot=1")
			if err != nil {
				return
			}
			stmt.Exec(vote.Message, clms.UserID)
			return
		} else {
			json.NewEncoder(w).Encode(Response{"invalid action", "invalid_action"})
			return
		}
	} else {
		// We need to check and limit their votes
		var count int
		count = -17
		rows, err := db.Query("select count(*) as count from votes where userid=? and message=?", clms.UserID, vote.Message)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		for rows.Next() {
			rows.Scan(&count)
		}
		fmt.Println("Count returned")
		fmt.Println(count)
		
		if count == 0 {
			// We can safely preform the action they want
			if vote.Doot == 1 {
				stmt, err := db.Prepare("INSERT votes SET message=?,userid=?,updoot=1")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
				return
			} else if vote.Doot == 0 {
				fmt.Println("applied downdoot")
				stmt, err := db.Prepare("INSERT votes SET message=?,userid=?,downdoot=1")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
				return
			} else {
				json.NewEncoder(w).Encode(Response{"invalid action", "invalid_action"})
				return
			}
		} else if count == 1{
			// otherwise we simply update the entry that already exists
			if vote.Doot == 1 {
				stmt, err := db.Prepare("UPDATE votes SET updoot=1, downdoot=0 where message=? and userid=?")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
				return
			} else if vote.Doot == 0 {
				stmt, err := db.Prepare("UPDATE votes SET updoot=0, downdoot=1 where message=? and userid=?")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
				return
			} else {
				json.NewEncoder(w).Encode(Response{"invalid action", "invalid_action"})
				return
			}
		} else {
			json.NewEncoder(w).Encode(Response{"something bad happened", "invalid_action"})
			return
		}
	}
}

func Register(w http.ResponseWriter, req *http.Request) {
	var tr TokenRequest

	json.NewDecoder(req.Body).Decode(&tr)

	if tr.UserID == "" {
		return
	}
	stmt, err := db.Prepare("INSERT users SET userid=?")
	if err != nil {
		return
	}
	stmt.Exec(tr.UserID)
}