package env

var (
	GithubRepositoryOwner   = getEnv("GITHUB_OWNER")
	GithubRepository        = getEnv("GITHUB_REPO")
	GithubPullRequestNumber = getEnv("GITHUB_PR_NUMBER")
	GithubAPIToken          = getEnv("PR_COMMENTATOR_GITHUB_API_TOKEN")
)
