package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content`
}

type Articles [2]Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{
			Title:   "Sample Title",
			Desc:    "Sample Description",
			Content: "Sample Content",
		},
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Test Content",
		},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode((articles))
}

func saveArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Save Article Endpoint Hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", saveArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8181", myRouter))
}

func main() {
	handleRequest()
}
