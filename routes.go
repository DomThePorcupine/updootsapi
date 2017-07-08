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
		"GetMessages",
		"GET",
		"/message",
		GetAllMessages,
	},
	Route{
		"CreateMessage",
		"POST",
		"/message",
		CreateMessage,
	},
}

var routes = Routes{
	Route{
		"GetToken",
		"POST",
		"/token",
		GetToken,
	},
}
