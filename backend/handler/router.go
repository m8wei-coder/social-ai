package handler

import (
    "net/http" 

    "github.com/gorilla/mux"   
)

func InitRouter() *mux.Router {
    router := mux.NewRouter()
    router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST")
    router.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")

    router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
    router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")
    
    return router
}