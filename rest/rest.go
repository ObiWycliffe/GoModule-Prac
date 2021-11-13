package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func main() {
	handleRequest()
}
