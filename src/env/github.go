package env

var (
	GithubRepositoryOwner   = getEnv("OWNER")
	GithubRepository        = getEnv("REPO")
	GithubPullRequestNumber = getEnv("PULL_NUMBER")
	GithubAPIToken          = getEnv("PR_COMMENTATOR_GITHUB_API_TOKEN")
)
