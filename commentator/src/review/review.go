//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Reviewer
package review

import (
	"context"
	"fmt"
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
	Text      string

	CustomCommentText *string // flag値としてユーザーが設定するコメント用のフォーマット
}

func (c Content) Message() string {
	var text string
	if c.CustomCommentText != nil {
		text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *c.CustomCommentText)
	} else {
		text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", c.FilePath, c.LineNum, c.Linter, c.Text) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
	}
	return text
}
