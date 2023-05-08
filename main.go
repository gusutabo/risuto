package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Words struct {
	ID      string `string:"id,omitempty"`
	Word    string `string:"word,omitempty"`
	Meaning string `string:"meaning,omitempty"`
	Status  string `string:"status,omitempty"`
}

var words []Words

func GetWords(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(words)
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range words {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateWord(w http.ResponseWriter, r *http.Request) {
	var word Words
	_ = json.NewDecoder(r.Body).Decode(&word)
	words = append(words, word)
	json.NewEncoder(w).Encode(words)
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range words {
		if item.ID == params["id"] {
			words = append(words[:index], words[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(words)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetWords).Methods("GET")
	router.HandleFunc("/{id}", GetWord).Methods("GET")
	router.HandleFunc("/word", CreateWord).Methods("POST")
	router.HandleFunc("/word/{id}", DeleteWord).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
