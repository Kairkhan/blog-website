package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(posts)
}

func store(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	posts = append(posts, post)
	json.NewEncoder(writer).Encode(post)
}

func view(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

type Post struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Category    []Category `json:"categories"`
}

type Category struct {
	Name string `json:"name"`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./frontend"))
	router.Handle("/", fileServer)
	router.HandleFunc("/posts", index).Methods("GET")
	router.HandleFunc("/posts/{id}", view).Methods("GET")
	router.HandleFunc("/posts", store).Methods("POST")

	fmt.Printf("Server starting on port 8001\n")

	log.Fatal(http.ListenAndServe(":8001", router))
}
