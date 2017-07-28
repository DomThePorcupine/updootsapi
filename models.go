package main
import "time"
type VoteResponse struct {
	Message string	`json:message,omitempty`
	Updoots int		`json:updoots,omitempty`
}

/*
Vote is a struct
*/
type Vote struct {
	Message int 	`json:message_id,omitempty`
	UserID	string	`json:userid,omitempty`
	Doot	int		`json:doot,omitempty`
}
/*
Response is a struct
*/
type Response struct {
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
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
MessageResponse is a struct
*/
type MessageResponse struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  string `json:"userid,omitempty"`
	Updoots int    `json:"updoots"`
	Vote	int	   `json:"vote"`
	Time	time.Time `json:"time"`
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

/*
TokenRequest is a struct
*/
type TokenRequest struct {
	UserID string  `json:"userid"`
}