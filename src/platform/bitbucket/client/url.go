package client

import (
	"fmt"

	"github.com/zawa-t/pr/reporter/platform/http/url"

	"github.com/zawa-t/pr/reporter/env"
)

var baseURL = fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s", env.Bitbucket.WorkspaceName, env.Bitbucket.RepositoryName)

var (
	prCommentPath = fmt.Sprintf("/pullrequests/%s/comments", env.Bitbucket.PRID)
	reportPath    = fmt.Sprintf("/commit/%s/reports", env.Bitbucket.CommitID)
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
