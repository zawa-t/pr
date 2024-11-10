package local

import (
	"context"
	"fmt"

	"github.com/zawa-t/pr/commentator/src/log"
	"github.com/zawa-t/pr/commentator/src/platform"
)

// Review ...
type Review struct {
}

// NewReview ...
func NewReview() *Review {
	return &Review{}
}

// AddComments ...
func (r *Review) AddComments(ctx context.Context, input platform.Data) error {
	for i, content := range input.Contents {
		var text string
		if content.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *content.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", content.FilePath, content.LineNum, content.Linter, content.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}
		input.Contents[i].Message = text
	}
	log.PrintJSON("input", input)
	return nil
}
