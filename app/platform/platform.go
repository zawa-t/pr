//go:generate moq -rm -out $GOPATH/app/test/mock/$GOFILE -pkg mock . Review
package platform

import (
	"context"
	"errors"
)

type PullRequest struct {
	Review
}

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
	Summary           string
	Details           string
	CustomCommentText *string // flag値としてユーザーが設定するコメント用のフォーマット
}

var ErrNotFound = errors.New("not found")

const (
	Bitbucket = "bitbucket"
	Github    = "github"
)
