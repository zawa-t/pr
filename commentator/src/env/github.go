package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type GithubConfig struct {
	RepositoryName    string
	PullRequestNumber string
	CommitID          string
	APIToken          string
}

var Github GithubConfig

func init() {
	if os.Getenv("GITHUB_ACTIONS") == "true" || Env.IsLocal() {
		Github = GithubConfig{
			RepositoryName: getEnv("GITHUB_REPOSITORY"),
			APIToken:       getEnv("PR_COMMENTATOR_GITHUB_API_TOKEN"),
		}

		Github.PullRequestNumber, Github.CommitID = getGithubPRNumber()
	}
}

type Head struct {
	SHA string `json:"sha"`
}

type PullRequest struct {
	Number int  `json:"number"`
	Head   Head `json:"head"`
}

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
}

// HACK: 以下のやり方はもう少し検討
func getGithubPRNumber() (string, string) {
	if Env.IsLocal() {
		return getEnv("GITHUB_PR_NUMBER"), getEnv("GITHUB_SHA")
	} else {
		data, err := os.ReadFile(getEnv("GITHUB_EVENT_PATH"))
		if err != nil {
			fmt.Printf("Error reading event file: %v\n", err)
			return "", ""
		}

		var event Event
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return "", ""
		}

		return strconv.Itoa(event.PullRequest.Number), event.PullRequest.Head.SHA
	}
}
