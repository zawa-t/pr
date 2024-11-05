package env

import (
	"log/slog"
	"strconv"
)

var (
	BitbucketWorkspace      = getEnv("WORKSPACE")
	BitbucketPRID           = getPRID()
	BitbucketRepositoryName = getEnv("REPOSITORY_NAME")
	BitbucketUserName       = getEnv("BITBUCKET_USERNAME")
	BitbucketAppPassword    = getEnv("BITBUCKET_APP_PASSWORD")
	BitbucketCommit         = getEnv("BITBUCKET_COMMIT")
)

func getPRID() int {
	prIDstr := getEnv("BITBUCKET_PR_ID")
	prID, err := strconv.Atoi(prIDstr)
	if err != nil {
		slog.Warn("Failed to exec strconv.Atoi().", "BITBUCKET_PR_ID", prIDstr)
	}
	return prID
}
