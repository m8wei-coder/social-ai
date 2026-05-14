package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialai/model"
	"socialai/service"
)

// use the pointer to save memory and avoid unnecessary copying of data
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Process request: Decode the JSON body into the struct

	fmt.Println("Received one upload request")
    decoder := json.NewDecoder(r.Body)
    var p model.Post
    if err := decoder.Decode(&p); err != nil {
        panic(err)
    }

	// 2. Business logic
	// TODO

	// 3. Respond to the client
	fmt.Fprintf(w, "Post received: %s\n", p.Message)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Received one request for search")
   w.Header().Set("Content-Type", "application/json")

   // 1. process the request
   // URL param -> string variable
   user := r.URL.Query().Get("user")
   keywords := r.URL.Query().Get("keywords")

   // 2. call service to handle request
   var posts []model.Post
   var err error

   // TODO: if both user and keywords are provided, search by user and keywords (Boolean query)

   // if user is not empty, search by user, otherwise search by keywords
   if user != "" {
       posts, err = service.SearchPostsByUser(user)
   } else {
       posts, err = service.SearchPostsByKeywords(keywords)
   }
   if err != nil {
       http.Error(w, "Failed to read post from backend", http.StatusInternalServerError)
       fmt.Printf("Failed to read post from backend %v.\n", err)
       return
   }

   // 3. construct response
   js, err := json.Marshal(posts)
   if err != nil {
       http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
       fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
       return
   }
   w.Write(js)
}