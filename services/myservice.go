package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Space struct {
	ID     string `json:"id"`
	Booked bool   `json:"booked"`
}

var spaces []Space

func GetSpacesEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spaces)
}

func GetSpaceEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreateSpaceEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeleteSpaceEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/spaces", GetSpacesEndpoint).Methods("GET")
	router.HandleFunc("/spaces/{id}", GetSpaceEndpoint).Methods("GET")
	router.HandleFunc("/spaces/{id}", CreateSpaceEndpoint).Methods("POST")
	router.HandleFunc("/spaces/{id}", DeleteSpaceEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func init() {
	spaces = append(spaces, Space{ID: "1", Booked: false})
	spaces = append(spaces, Space{ID: "2", Booked: true})
	spaces = append(spaces, Space{ID: "3", Booked: false})
}
