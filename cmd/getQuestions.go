/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/spf13/cobra"
    "fastrack/quiz/backend"
    "bufio"
    "os"
    "strings"
    "bytes"
    "log"
)

var getQuestionsCmd = &cobra.Command{
	Use:   "getQuestions",
	Short: "Start quiz",
	Long: `The quiz will start and wait the users inputs every answer after
	every given question. Before finishing, it will return the amount of correct
	answers and the times better than other users in percentage. After this the
	command will finish and it will have to be run again to play the same game again.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getQuestions called")
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)

    // Get questions from backend
    body, err := fetchQuestions("http://localhost:8080/questions")
    handleError(err, "Error:")

    questions, err := unmarshalQuestions(body)
    handleError(err, "Error unmarshalling response:")


    // Display questions and read user answers.
    var answers []backend.Answer

    for _, q := range questions {
        fmt.Printf("%d) %s \n", q.Id, q.Question)

        for i, option := range q.Options {
            fmt.Printf("\t %d) %s \n", i+1, option)
        }

        answers = readAnswerFromInput(answers, q)
    }

    jsonData, err := json.Marshal(answers)
    handleError(err, "Error marshalling answers:")

    // submit answers to backend
    rawResp, err := http.Post("http://localhost:8080/submit-answers", "application/json", bytes.NewBuffer(jsonData))

    if rawResp.StatusCode >= 400 {
        handleError(fmt.Errorf("server returned error http status: %d", rawResp.StatusCode), "Error:")
        return
    }

    handleError(err, "Error posting the answers")

    resp, err := ioutil.ReadAll(rawResp.Body)
    handleError(err, "Error reading response body after submitting answers")

    var result map[string]interface{}
    json.Unmarshal(resp, &result)

    fmt.Printf("You got %d correct answers! You are better than %d%% of users!\n", int(result["correctAnswers"].(float64)), int(result["percentageBetterThanOtherUsers"].(float64)))
}

func readAnswerFromInput(answers []backend.Answer, q backend.Question) []backend.Answer {
    reader := bufio.NewReader(os.Stdin)

    input, _ := reader.ReadString('\n')

    input = strings.TrimSpace(input)

    answer := 0

    fmt.Sscanf(input, "%d", &answer)

    return append(answers, backend.Answer{
        QuestionId: q.Id,
        Answer: answer - 1,
    })
}

func fetchQuestions(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func unmarshalQuestions(body []byte) ([]backend.Question, error) {
    var questions []backend.Question
    err := json.Unmarshal(body, &questions)
    if err != nil {
        return nil, err
    }
    return questions, nil
}

func handleError(err error, message string) {
    if err != nil {
        log.Printf("%s: %v\n", message, err)
        os.Exit(1)
    }
}
