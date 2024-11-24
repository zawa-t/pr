package role

const (
	LocalComment = iota
	BitbucketPRComment
	GithubPRComment
	// GithubPRCheck
	GithubCheck
	// GithubPRReviewer
	maxRole = iota // MEMO: 定義している定数の数を計測するための定数（iotaは「最後の定数の次に設定されている値」という点を明示するため設定）
)

var NameList = map[string]int{
	"local-comment":        LocalComment,
	"bitbucket-pr-comment": BitbucketPRComment,
	"github-pr-comment":    GithubPRComment,
	// "github-pr-check":      GithubPRCheck,
	"github-check": GithubCheck,
	// "github-pr-reviewer":      GithubPRReviewer,
}
