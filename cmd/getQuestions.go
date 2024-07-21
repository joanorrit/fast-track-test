/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getQuestionsCmd represents the getQuestions command
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

	questions := map[string]string{
	   "1" : "What is the capital of Spain?",
	   "2" : "How much are 2 + 2?",
	}

	questionOptions := map[string][]string{
	    "1" : {"Madrid", "Barcelona", "Seville", "Valencia"},
	    "2" : {"2", "4", "5", "6"},
    }

// CHECK. Sort map so index are shown in order.
    for questionId, question := range questions {
        fmt.Println(questionId + ") " + question)
        optionIndex := "a"
        for _, option := range questionOptions[questionId] {
            fmt.Println("\t" + optionIndex + ") " + option)
            optionIndex = string(rune(optionIndex[0]) + 1)
        }
    }
}
