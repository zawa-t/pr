//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Reviewer
package review

import (
	"context"
)

// Reviewer ...
type Reviewer interface {
	Review(ctx context.Context, data Data) error
}

// Data ...
type Data struct {
	Name     string
	Contents []Content
}

// Content ...
type Content struct {
	Linter    string
	FilePath  string
	LineNum   uint
	ColumnNum uint
	CodeLine  string
	Indicator string
	Message   string

	CustomCommentText *string // flag値としてユーザーが設定するコメント用のフォーマット
}
