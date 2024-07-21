package backend

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
)

type Question struct {
    Id       int      `json:"id"`
    Question string   `json:"question"`
    Options  []string `json:"answers"`
    Correct  int      `json:"-"`
}

type Answer struct {
    QuestionId int `json:"question_id"`
    AnswerId     int `json:"answer"`
}

var questions = []Question{
    {Id: 1, Question: "What is the capital of Spain?", Options: []string{"Madrid", "Barcelona", "Seville", "Valencia"}, Correct: 0},
    {Id: 2, Question: "How much are 2 + 2?", Options: []string{"4", "2", "5", "6"}, Correct: 1},
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(questions)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    var answers []Answer
    err = json.Unmarshal(body, &answers)
    if err != nil {
        http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(answers)
}

func StartBackend() {
    r := mux.NewRouter()
    r.HandleFunc("/questions", GetQuestions).Methods("GET")
    r.HandleFunc("/submit-answers", SubmitAnswers).Methods("POST")

    http.ListenAndServe(":8080", r)
}