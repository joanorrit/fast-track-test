package backend

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Question struct {
    Id       int      `json:"id"`
    Question string   `json:"question"`
    Options  []string `json:"answers"`
    Correct  int      `json:"-"`
}

var questions = []Question{
    {Id: 1, Question: "What is the capital of Spain?", Options: []string{"Madrid", "Barcelona", "Seville", "Valencia"}, Correct: 0},
    {Id: 2, Question: "How much are 2 + 2?", Options: []string{"4", "2", "5", "6"}, Correct: 1},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(questions)

}

func StartBackend() {
    r := mux.NewRouter()
    r.HandleFunc("/questions", HomeHandler)

    http.ListenAndServe(":8080", r)
}