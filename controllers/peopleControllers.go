package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"awesomeProject/models"
)

var people []models.Person

var GetPeopleEndpoint = func(w http.ResponseWriter, r *http.Request) {
	log.Println("GetPeopleEndpoint")
	json.NewEncoder(w).Encode(people)
}
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Println("GetPersonEndpoint {}", params["id"])

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Person{})
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
func init() {
	people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
	people = append(people, models.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &models.Address{City: "City Z", State: "State Y"}})

}
