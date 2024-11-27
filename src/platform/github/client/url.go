package client

import (
	"fmt"

	"github.com/zawa-t/pr/src/platform/http/url"

	"github.com/zawa-t/pr/src/env"
)

var baseURL = fmt.Sprintf("https://api.github.com/repos/%s", env.Github.RepositoryName)

var (
	prCommentPath = fmt.Sprintf("/pulls/%s/comments", env.Github.PullRequestNumber)
	prReviewPath  = fmt.Sprintf("/pulls/%s/reviews", env.Github.PullRequestNumber)
	checkRunPath  = "/check-runs"
)

var (
	prCommentURL = url.JoinPathWithNoError(baseURL, prCommentPath)
	prReviewURL  = url.JoinPathWithNoError(baseURL, prReviewPath)
	checkRunURL  = url.JoinPathWithNoError(baseURL, checkRunPath)
)
