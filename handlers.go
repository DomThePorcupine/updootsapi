package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

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
	
	rows, err := db.Query("SELECT admin, userid from users where userid = ?", tr.UserID)

	if err != nil {
		// Internal server error
		fmt.Println(err.Error())
		return
	}

	var truefalse int
	var uid string
	truefalse = -1
	// Grab the value we want
	for rows.Next() {
		rows.Scan(&truefalse, &uid)
	}

	if(truefalse == -1) {
		json.NewEncoder(w).Encode(Response{"No user id given", "invalid_id"})
		return
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
	claims["userid"] = uid
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
func GetAllMessagesNew(w http.ResponseWriter, req *http.Request) {
	// Get our claims we will need the user id
	clms, ok := req.Context().Value(Claims{}).(Claims)
	if !ok {
		json.NewEncoder(w).Encode(Response{"invalid id", "invalid_id"})
		return
	}

	rows, err := db.Query("select messages.created, id, messages.message, ifnull(doots,0) as totalvotes " + 
							"from messages left join(select votes.message, " +
							"cast((sum(votes.updoot) - sum(votes.downdoot)) as signed) " + 
							"as doots from votes group by votes.message) as votes " + 
							"on messages.id = votes.message having totalvotes > -3 order by messages.created desc")
							
	// Need to also get what we voted on for visual ques
	votedrows, err := db.Query("select updoot, downdoot, message from votes where userid=?", clms.UserID)

	// Create a hash this will be usefull later
	votes := make(map[int]int)
	var up int
	var down int
	var id int
	
	for votedrows.Next() {
		votedrows.Scan(&up, &down, &id)
		if(up == 1) {
			votes[id] = 1
		} else if(down == 1) {
			votes[id] = -1
		}
		// In go if an int doesn't
		// exist in a map it is simply zerp
		// so we can just let that do the hard work
	}
	// If we experience some kind of error
	if err != nil {
		checkError(err)
		w.WriteHeader(500)
		w.Write([]byte("Uh oh!"))
		return
	}

	var mess string
	var ups int
	var tim time.Time
	ups = -17
	// These are the messages we will
	// be sending back
	var messages []MessageResponse
	//fmt.Println(rows)
	for rows.Next() {
		var message MessageResponse

		rows.Scan(&tim, &id, &mess, &ups)
		if ups == -17 {
			message.Updoots = 0
		} else {
			message.Updoots = ups
		}
		message.ID = id
		//fmt.Println(mess)
		message.Message = mess
		message.Vote = votes[id]
		message.Time = tim
		messages = append(messages, message)
		ups = -17
	}

	// Makes sure the client sees application/json
	w.WriteHeader(http.StatusOK)
	if messages == nil {
		messages = make([]MessageResponse, 0)
	}

	// else we should have our rows
	json.NewEncoder(w).Encode(messages)
}

/*
GetAllMessages is a function
*/
func GetAllMessagesTop(w http.ResponseWriter, req *http.Request) {
	// Get our claims we will need the user id
	clms, ok := req.Context().Value(Claims{}).(Claims)
	if !ok {
		json.NewEncoder(w).Encode(Response{"invalid id", "invalid_id"})
		return
	}

	rows, err := db.Query("select messages.created, id, messages.message, ifnull(doots,0) as totalvotes " + 
							"from messages left join(select votes.message, " +
							"cast((sum(votes.updoot) - sum(votes.downdoot)) as signed) " + 
							"as doots from votes group by votes.message) as votes " + 
							"on messages.id = votes.message having totalvotes > -3 order by ifnull(doots,0) desc, messages.created desc")

	// Need to also get what we voted on for visual ques
	votedrows, err := db.Query("select updoot, downdoot, message from votes where userid=?", clms.UserID)

	// Create a hash this will be usefull later
	votes := make(map[int]int)
	var up int
	var down int
	var id int
	
	for votedrows.Next() {
		votedrows.Scan(&up, &down, &id)
		if(up == 1) {
			votes[id] = 1
		} else if(down == 1) {
			votes[id] = -1
		}
		// In go if an int doesn't
		// exist in a map it is simply zerp
		// so we can just let that do the hard work
	}
	// If we experience some kind of error
	if err != nil {
		checkError(err)
		w.WriteHeader(500)
		w.Write([]byte("Uh oh!"))
		return
	}

	var mess string
	var ups int
	var tim time.Time
	ups = -17
	// These are the messages we will
	// be sending back
	var messages []MessageResponse
	//fmt.Println(rows)
	for rows.Next() {
		var message MessageResponse

		rows.Scan(&tim, &id, &mess, &ups)
		if ups == -17 {
			message.Updoots = 0
		} else {
			message.Updoots = ups
		}
		message.ID = id
		//fmt.Println(mess)
		message.Message = mess
		message.Vote = votes[id]
		message.Time = tim
		messages = append(messages, message)
		ups = -17
	}

	// Makes sure the client sees application/json
	w.WriteHeader(http.StatusOK)
	if messages == nil {
		messages = make([]MessageResponse, 0)
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
	fmt.Println(string(nMessage.Message))
	fmt.Println(utf8.ValidString(nMessage.Message))

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
	stmt, err := db.Prepare("INSERT into messages (message, userid) values( ? , ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	_, err = stmt.Exec(nMessage.Message, clms.UserID)

	if err != nil {
		checkError(err)
	}
	//id, err := res.LastInsertId()
	
	// return the newly created object
	json.NewEncoder(w).Encode(Empty{})
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
			} else if vote.Doot == 0 {
				fmt.Println("applied downdoot")
				stmt, err := db.Prepare("INSERT votes SET message=?,userid=?,downdoot=1")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
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
			} else if vote.Doot == 0 {
				stmt, err := db.Prepare("UPDATE votes SET updoot=0, downdoot=1 where message=? and userid=?")
				if err != nil {
					fmt.Println("BADDDDDDDDDDD")
					return
				}
				stmt.Exec(vote.Message, clms.UserID)
			} else {
				json.NewEncoder(w).Encode(Response{"invalid action", "invalid_action"})
				return
			}
		} else {
			json.NewEncoder(w).Encode(Response{"something bad happened", "invalid_action"})
			return
		}
	}

	rows, err := db.Query("select ifnull((sum(votes.updoot) - sum(votes.downdoot)),0) as updoots from votes where votes.message=?", vote.Message)
	if err != nil {
		fmt.Println("BADDDDDDDDDDD")
		json.NewEncoder(w).Encode(Response{"something bad happened", "invalid_action"})
		return
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	json.NewEncoder(w).Encode(VoteResponse{"success", count})
	return
}

func Register(w http.ResponseWriter, req *http.Request) {
	var tr TokenRequest

	json.NewDecoder(req.Body).Decode(&tr)

	if tr.UserID == "" {
		return
	}
	
	fmt.Println(tr.UserID)

	stmt, err := db.Prepare("INSERT into users (userid) values(?)")
	if err != nil {
		json.NewEncoder(w).Encode(Response{"something bad happened first", "invalid_action"})
		return
	}
	_, err = stmt.Exec(tr.UserID)

	if err != nil {
		json.NewEncoder(w).Encode(Response{"something bad happened", "invalid_action"})
		return
	}

}