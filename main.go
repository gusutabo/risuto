package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

type Word struct {
    ID       string        `json:"id,omitempty"`
    Word     string        `json:"word,omitempty"`
    Meaning  string        `json:"meaning,omitempty"`
}

var word []Word

func getWord(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    for _, item := range word {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    json.NewEncoder(w).Encode(word)
}

func createWord(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var new_word Word
    _ = json.NewDecoder(r.Body).Decode(&new_word)
    new_word.ID = params["id"]
    word = append(word, new_word)

    json.NewEncoder(w).Encode(word)
}

func delWord(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    for index, item := range word {
        if item.ID == params["id"] {
            word = append(word[:index], word[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(word)
    }
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/words", getWord).Methods("GET")
    router.HandleFunc("/words/{id}", getWord).Methods("GET")
    router.HandleFunc("/words/{id}", createWord).Methods("POST")
    router.HandleFunc("/words/{id}", delWord).Methods("DELETE")

    http.ListenAndServe(":8000", router)
}