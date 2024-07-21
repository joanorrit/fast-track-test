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
)

var getQuestionsCmd = &cobra.Command{
	Use:   "getQuestions",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getQuestions called")
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)

    body, err := fetchQuestions("http://localhost:8080/questions")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    questions, err := unmarshalQuestions(body)
    if err != nil {
        fmt.Println("Error unmarshalling response:", err)
        return
    }

    var answers []backend.Answer

    for _, q := range questions {
        fmt.Printf("%d) %s \n", q.Id, q.Question)

        for i, option := range q.Options {
            fmt.Printf("\t %d) %s \n", i+1, option)
        }

        answers = readAnswerFromInput(answers, q)
    }

    jsonData, err := json.Marshal(answers)
    if err != nil {
        fmt.Println("Error marshalling answers:", err)
        return
    }
    fmt.Println(answers)
    resp, err := http.Post("http://localhost:8080/submit-answers", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    fmt.Println(resp)
}

func readAnswerFromInput(answers []backend.Answer, q backend.Question) []backend.Answer {
    reader := bufio.NewReader(os.Stdin)

    input, _ := reader.ReadString('\n')

    input = strings.TrimSpace(input)

    answer := 0

    fmt.Sscanf(input, "%d", &answer)

    return append(answers, backend.Answer{
        QuestionId: q.Id,
        AnswerId: answer - 1,
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

func displayQuestions(questions []backend.Question) {
    for _, q := range questions {
        fmt.Printf("%d) %s \n", q.Id, q.Question)
        for i, option := range q.Options {
            fmt.Printf("\t %d) %s \n", i+1, option)
        }
    }
}
