package role

const (
	LocalComment = iota
	BitbucketPRComment
	GithubPRComment
	// GithubPRCheck
	GithubCheck
)

var NameList = map[string]int{
	"local-comment":        LocalComment,
	"bitbucket-pr-comment": BitbucketPRComment,
	"github-pr-comment":    GithubPRComment,
	// "github-pr-check":      GithubPRCheck,
	"github-check": GithubCheck,
}
