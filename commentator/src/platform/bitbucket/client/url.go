package client

import (
	"fmt"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/platform/http/url"
)

var baseURL = fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s", env.BitbucketWorkspace, env.BitbucketRepositoryName)

var (
	prCommentPath = fmt.Sprintf("/pullrequests/%d/comments", env.BitbucketPRID)
	reportPath    = fmt.Sprintf("/commit/%s/reports", env.BitbucketCommit)
)

var (
	prCommentURL = url.JoinPathWithNoError(baseURL, prCommentPath)
	reportURL    = func(reportID string) string {
		return url.JoinPathWithNoError(baseURL, reportPath, reportID)
	}
	bulkAnnotationsURL = func(reportID string) string {
		return url.JoinPathWithNoError(baseURL, reportPath, reportID, "/annotations")
	}
)
