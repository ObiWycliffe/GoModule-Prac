package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json: "id, omitempty"`
	Firstname string   `json: "firstname, omitempty"`
	Lastname  string   `json: "lastname, omitempty"`
	Address   *Address `json: "address, omitempty"`
}

type Address struct {
	City  string `json: "city, omitempty"`
	State string `json: "state, omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}

func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["ID"] {
			people = append(people[:index], people[index+1:]...) // append people before and after index being deletedss
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", Firstname: "Obi", Lastname: "Wycliffe", Address: &Address{City: "Nairobi", State: "Kenya"}})
	people = append(people, Person{ID: "2", Firstname: "Ken", Lastname: "Jaramogi"})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{ID}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8181", router))
}
