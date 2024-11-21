package env

import (
	"os"
)

type BitbucketConfig struct {
	WorkspaceName  string
	PRID           string
	RepositoryName string
	UserName       string
	AppPassword    string
	CommitID       string
}

var Bitbucket BitbucketConfig

func init() {
	if os.Getenv("BITBUCKET_BUILD_NUMBER") != "" || Env.IsLocal() {
		Bitbucket = BitbucketConfig{
			WorkspaceName:  getEnv("WORKSPACE"),
			PRID:           getEnv("BITBUCKET_PR_ID"),
			RepositoryName: getEnv("REPOSITORY_NAME"),
			UserName:       getEnv("BITBUCKET_USERNAME"),
			AppPassword:    getEnv("BITBUCKET_APP_PASSWORD"),
			CommitID:       getEnv("BITBUCKET_COMMIT"),
		}
	}
}
