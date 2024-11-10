package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var (
	GithubRepository        = getEnv("GITHUB_REPOSITORY")
	GithubPullRequestNumber = getGithubPRNumber()
	GithubAPIToken          = getEnv("PR_COMMENTATOR_GITHUB_API_TOKEN")
	GithubCommitID          = getEnv("GITHUB_SHA")
)

type PullRequest struct {
	Number int `json:"number"`
}

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
}

func getGithubPRNumber() string {
	if Env.IsLocal() {
		return getEnv("GITHUB_PR_NUMBER")
	} else {
		data, err := os.ReadFile(getEnv("GITHUB_EVENT_PATH"))
		if err != nil {
			fmt.Printf("Error reading event file: %v\n", err)
			return ""
		}

		var event Event
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return ""
		}

		fmt.Printf("Pull Request Number: %d\n", event.PullRequest.Number)
		return strconv.Itoa(event.PullRequest.Number)
	}
}
