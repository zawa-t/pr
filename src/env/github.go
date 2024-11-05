package env

var (
	GithubRepositoryOwner   = getEnv("OWNER")
	GithubRepository        = getEnv("REPO")
	GithubPullRequestNumber = getEnv("PULL_NUMBER")
)
