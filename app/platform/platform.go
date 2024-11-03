package platform

import (
	"context"
	"errors"
)

// PullRequest ...
type PullRequest struct {
	Comment
}

// NewPullRequest ...
func NewPullRequest(c Comment) *PullRequest {
	return &PullRequest{c}
}

// Comment ...
type Comment interface {
	AddComments(ctx context.Context, data Input) error
}

// Input ...
type Input struct {
	Name  string // nameが指定されていない場合はLinterを設定？
	Datas []Data
}

// Data ...
type Data struct {
	Linter            string
	FilePath          string
	LineNum           uint
	Summary           string
	Details           string
	CustomCommentText *string
}

var ErrNotFound = errors.New("not found")
