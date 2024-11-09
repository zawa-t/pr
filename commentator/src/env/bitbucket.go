package env

import (
	"log/slog"
	"strconv"
)

var (
	BitbucketWorkspace      = getEnv("WORKSPACE")
	BitbucketPRID           = getBitbucketPRID("BITBUCKET_PR_ID")
	BitbucketRepositoryName = getEnv("REPOSITORY_NAME")
	BitbucketUserName       = getEnv("BITBUCKET_USERNAME")
	BitbucketAppPassword    = getEnv("BITBUCKET_APP_PASSWORD")
	BitbucketCommit         = getEnv("BITBUCKET_COMMIT")
)

func getBitbucketPRID(key string) int {
	prIDstr := getEnv(key)
	prID, err := strconv.Atoi(prIDstr)
	if err != nil {
		slog.Warn("The following value should be an integer, but it isn't.", "BITBUCKET_PR_ID", prIDstr)
	}
	return prID
}
