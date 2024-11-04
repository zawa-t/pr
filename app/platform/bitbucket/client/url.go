package client

import (
	"fmt"

	"github.com/zawa-t/pr-commentator/env"
)

var baseURL = fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s", env.BitbucketWorkspace, env.BitbucketRepositoryName)

var (
	prCommentPath   = fmt.Sprintf("/pullrequests/%d/comments", env.BitbucketPRID)
	reportPath      = fmt.Sprintf("/commit/%s/reports", env.BitbucketCommit)
	annotationsPath = "/annotations"
)

var (
	prCommentURL       = baseURL + prCommentPath
	reportURL          = func(reportID string) string { return baseURL + reportPath + "/" + reportID }
	bulkAnnotationsURL = func(reportID string) string { return baseURL + reportPath + "/" + reportID + annotationsPath }
)
