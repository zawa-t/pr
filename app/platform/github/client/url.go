package client

import (
	"fmt"

	"github.com/zawa-t/pr-commentator/env"
)

var baseURL = fmt.Sprintf("https://api.github.com/repos/%s/%s", env.GithubRepositoryOwner, env.GithubRepository)

var (
	prCommentPath = fmt.Sprintf("/pulls/%s/comments", env.GithubPullRequestNumber)
)

var (
	prCommentURL = baseURL + prCommentPath
)
