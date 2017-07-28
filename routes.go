package main

import "net/http"

/*
Route is a struct
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*
Routes is a slice of Route structs
*/
type Routes []Route

var protectedroutes = Routes{
	Route{
		"GetMessagesNew",
		"GET",
		"/message/new",
		GetAllMessagesNew,
	},
	Route{
		"GetMessagesTop",
		"GET",
		"/message/top",
		GetAllMessagesTop,
	},
	Route{
		"CreateMessage",
		"POST",
		"/message",
		CreateMessage,
	},
	Route{
		"DootOnMessage",
		"POST",
		"/doot",
		DootOnMessage,
	},
}

var routes = Routes{
	Route{
		"GetToken",
		"POST",
		"/token",
		GetToken,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
}
