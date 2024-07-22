# fast-track-test

## How to run the quiz?
- Clone the repository and enter the directory with 
`cd {path-to-where-repository-lives}/fast-track-test`

- In the same terminal start the backend with
`go run startBackend.go`

- Open a separate terminal and enter the directory of the repository
`cd {path-to-where-repository-lives}/fast-track-test`

- Run the next cobra command to start the quiz:
`go run main.go getQuestions`

- The quiz will finish after you've given all the answers. You can start the quiz as many times as you want with the same command. The backend will keep track of all answers given in previous attempts and give you how many times are you better than other users in percentage.
