package client

import (
	"fmt"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/platform/http/url"
)

var baseURL = fmt.Sprintf("https://api.github.com/repos/%s", env.GithubRepository)

var (
	prCommentPath = fmt.Sprintf("/pulls/%s/comments", env.GithubPullRequestNumber)
	prReviewPath  = fmt.Sprintf("/pulls/%s/reviews", env.GithubPullRequestNumber)
	checkRunPath  = "/check-runs"
)

var (
	prCommentURL = url.JoinPathWithNoError(baseURL, prCommentPath)
	prReviewURL  = url.JoinPathWithNoError(baseURL, prReviewPath)
	checkRunURL  = url.JoinPathWithNoError(baseURL, checkRunPath)
)
