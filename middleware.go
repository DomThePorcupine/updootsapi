package main

import (
	"fmt"
	"net/http"
)

func validate(protectedPage http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Auth")
		if err != nil {
			http.NotFound(res, req)
			return
		}
		fmt.Println(cookie)
	})
}
