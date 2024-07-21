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

    resp, err := http.Get("http://localhost:8080/questions")

    if err != nil {
        fmt.Println("Error fetching questions:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    var questions []backend.Question
    err = json.Unmarshal(body, &questions)
    if err != nil {
        fmt.Println("Error unmarshalling response:", err)
        return
    }

    for _, q := range questions {
        fmt.Printf("%d) %s \n", q.Id, q.Question)
        optionIndex := 1
        for _, option := range q.Options {
            fmt.Printf("\t %d) %s \n", optionIndex, option)
            optionIndex++
        }
    }
}
