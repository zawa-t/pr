package platform

import (
	"context"
	"errors"
)

type Platform struct {
	PullRequest PullRequest
}

func New(pr PullRequest) *Platform {
	return &Platform{pr}
}

// PullRequest ...
type PullRequest interface {
	AddComments(ctx context.Context, data Input) error
}

// Input ...
type Input struct {
	Name  string
	Datas []Data
}

// Data ...
type Data struct {
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
