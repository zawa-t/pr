//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Review
package platform

import (
	"context"
	"errors"
)

// PullRequest ...
type PullRequest struct {
	Review
}

// NewPullRequest ...
func NewPullRequest(r Review) *PullRequest {
	return &PullRequest{r}
}

// Review ...
type Review interface {
	AddComments(ctx context.Context, data Data) error
}

// Data ...
type Data struct {
	Name     string
	RawDatas []Raw
}

// Raw ...
type Raw struct {
	Linter            string
	FilePath          string
	LineNum           uint
	Message           string
	CustomCommentText *string // flag値としてユーザーが設定するコメント用のフォーマット
}

var ErrNotFound = errors.New("not found")

const (
	Bitbucket = "bitbucket"
	Github    = "github"
)
