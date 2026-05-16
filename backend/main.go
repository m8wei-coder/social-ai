package main

import (
	"fmt"
	"log"
	"net/http"

	"socialai/handler"
	"socialai/repository"
)
func main() {
    fmt.Println("started-service")

    repository.InitElasticsearchBackend()
	repository.InitGCSBackend()

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}