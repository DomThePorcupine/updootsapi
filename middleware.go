package main

import (
	"net/http"
  "fmt"
)

func validate(protectedPage http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Authorization")
		if err != nil {
			http.NotFound(w, req)
			return
		}

    fmt.Println(cookie)
	  protectedPage.ServeHTTP(w, req)
	})
}
