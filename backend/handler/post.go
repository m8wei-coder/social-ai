package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialai/model"
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