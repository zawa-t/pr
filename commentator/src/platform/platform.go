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
	Contents []Content
}

// Content ...
type Content struct {
	Linter            string
	FilePath          string
	LineNum           uint
	ColumnNum         uint
	Message           string
	CustomCommentText *string // flag値としてユーザーが設定するコメント用のフォーマット
}

var ErrNotFound = errors.New("not found")

const (
	Local     = "local"
	Github    = "github"
	Bitbucket = "bitbucket"
)
