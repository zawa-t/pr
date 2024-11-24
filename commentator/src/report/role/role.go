package role

const (
	LocalComment = iota
	BitbucketPRComment
	GithubPRComment
	GithubPRCheck
	GithubCheck
)

// const (
// 	LocalComment       = "local-comment"
// 	BitbucketPRComment = "bitbucket-pr-comment"
// 	GithubPRComment    = "github-pr-comment"
// 	GithubPRCheck      = "github-pr-check"
// 	GithubCheck        = "github-check"
// )

// type roleName string

var NameList = map[string]int{
	"local-comment":        LocalComment,
	"bitbucket-pr-comment": BitbucketPRComment,
	"github-pr-comment":    GithubPRComment,
	"github-pr-check":      GithubPRCheck,
	"github-check":         GithubCheck,
}
