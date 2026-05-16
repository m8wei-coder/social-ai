package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"socialai/model"
	"socialai/service"

	"github.com/pborman/uuid"
)

var (
    mediaTypes = map[string]string{
       ".jpeg": "image",
       ".jpg":  "image",
       ".gif":  "image",
       ".png":  "image",
       ".mov":  "video",
       ".mp4":  "video",
       ".avi":  "video",
       ".flv":  "video",
       ".wmv":  "video",
    }
)

// use the pointer to save memory and avoid unnecessary copying of data
func uploadHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Received one upload request")
	// 1. Process request: Multipart text+file -> Post+file

    p := model.Post{
       Id:      uuid.New(),
       User:    r.FormValue("user"),
       Message: r.FormValue("message"),
    }

    file, header, err := r.FormFile("media_file")
    if err != nil {
       http.Error(w, "Media file is not available", http.StatusBadRequest)
       fmt.Printf("Media file is not available %v\n", err)
       return
    }

    // p.type, header->suffix->type
    suffix := filepath.Ext(header.Filename) // -> .jpg .avi .mp4
    if t, ok := mediaTypes[suffix]; ok {
        p.Type = t
    } else {
        p.Type = "unknown" // -> error out
    }

    err = service.SavePost(&p, file)
    if err != nil {
       http.Error(w, "Failed to save post to backend", http.StatusInternalServerError)
       fmt.Printf("Failed to save post to backend %v\n", err)
       return
    }

	// 3. construct response
	fmt.Println("Post is saved successfully.")

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