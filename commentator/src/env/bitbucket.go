package env

import (
	"log/slog"
	"os"
	"strconv"
)

type BitbucketConfig struct {
	Workspace      string
	PRID           int
	RepositoryName string
	UserName       string
	AppPassword    string
	Commit         string
}

var Bitbucket BitbucketConfig

func init() {
	if os.Getenv("BITBUCKET_BUILD_NUMBER") != "" {
		Bitbucket = BitbucketConfig{
			Workspace:      getEnv("WORKSPACE"),
			PRID:           getBitbucketPRID("BITBUCKET_PR_ID"),
			RepositoryName: getEnv("REPOSITORY_NAME"),
			UserName:       getEnv("BITBUCKET_USERNAME"),
			AppPassword:    getEnv("BITBUCKET_APP_PASSWORD"),
			Commit:         getEnv("BITBUCKET_COMMIT"),
		}

		// // Validate required fields
		// if Bitbucket.Workspace == "" || Bitbucket.RepositoryName == "" || Bitbucket.Commit == "" {
		// 	fmt.Println("Error: Missing required environment variables for Bitbucket")
		// 	os.Exit(1)
		// }
	}
}

func getBitbucketPRID(key string) int {
	prIDstr := getEnv(key)
	prID, err := strconv.Atoi(prIDstr)
	if err != nil {
		slog.Warn("The following value should be an integer, but it isn't.", "BITBUCKET_PR_ID", prIDstr)
	}
	return prID
}
