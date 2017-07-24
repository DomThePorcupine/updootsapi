package main

import (
	"net/http"
  "fmt"
	"context"
	jwt "github.com/dgrijalva/jwt-go"
)

func validate(protectedPage http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// If no Auth cookie is set then return a 404 not found
		cookie, err := req.Cookie("Authorization")
		if err != nil {
			http.NotFound(w, req)
			return
		}

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error){
				// Make sure token's signature wasn't changed
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected siging method")    
				}    
				return signingKey, nil
		})
		if err != nil {
				http.NotFound(w, req)
				return
		}
		
		// Grab the tokens claims and pass it into the original request
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), Claims{},  *claims)
			protectedPage.ServeHTTP(w, req.WithContext(ctx))
		} else {
				http.NotFound(w, req)
				return
		}
	  
	})
}
