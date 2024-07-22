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
    Answer   int `json:"answer"`
}

var questions = []Question{
    {Id: 1, Question: "What is the capital of Spain?", Options: []string{"Madrid", "Barcelona", "Seville", "Valencia"}, Correct: 0},
    {Id: 2, Question: "How much are 2 + 2?", Options: []string{"4", "2", "5", "6"}, Correct: 1},
    {Id: 3, Question: "Who wrote 'Romeo and Juliet'?", Options: []string{"William Shakespeare", "Charles Dickens", "Mark Twain", "Jane Austen"}, Correct: 0},
    {Id: 4, Question: "What is the largest planet in our solar system?", Options: []string{"Earth", "Mars", "Jupiter", "Saturn"}, Correct: 2},
    {Id: 5, Question: "What is the main ingredient in guacamole?", Options: []string{"Tomato", "Cucumber", "Avocado", "Lettuce"}, Correct: 2},
}

var userScores []int

func GetQuestions(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(questions)
}

func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
    // Get answers from body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }

    var answers []Answer
    err = json.Unmarshal(body, &answers)
    if err != nil {
        http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
        return
    }

    correctAnswers := getAmountOfCorrectAnswers(answers)

    percentageBetterThanOtherUsers := 0.0
    if len(userScores) > 0 {
        percentageBetterThanOtherUsers = getTimesBetterThanOtherUsers(correctAnswers)
    }

    // Update global user scores
    userScores = append(userScores, correctAnswers)

    json.NewEncoder(w).Encode(map[string]interface{}{
        "correctAnswers": correctAnswers,
        "percentageBetterThanOtherUsers" : percentageBetterThanOtherUsers,
    })
}

func GetQuestionByID(id int) (*Question) {
    for _, q := range questions {
        if q.Id == id {
            return &q
        }
    }
    return nil
}

func getAmountOfCorrectAnswers(answers []Answer) int {
    correctAnswers := 0
    for _, a := range answers {
        question := GetQuestionByID(a.QuestionId)
        if question.Correct == a.Answer {
            correctAnswers++
        }
    }
    return correctAnswers
}

func getTimesBetterThanOtherUsers(correctAnswers int) float64 {
    timesBetterOtherUsers := 0.0
    for _, previousUserScore := range userScores {
        if correctAnswers > previousUserScore {
            timesBetterOtherUsers++
        }
    }
    return float64(timesBetterOtherUsers) / float64(len(userScores)) * 100
}

func StartBackend() {
    r := mux.NewRouter()
    r.HandleFunc("/questions", GetQuestions).Methods("GET")
    r.HandleFunc("/submit-answers", SubmitAnswers).Methods("POST")

    http.ListenAndServe(":8080", r)
}