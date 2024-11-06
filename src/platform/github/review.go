package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/zawa-t/pr-commentator/src/env"
	"github.com/zawa-t/pr-commentator/src/log"
	"github.com/zawa-t/pr-commentator/src/platform"
)

// Review ...
type Review struct {
	client Client
}

// NewReview ...
func NewReview(c Client) *Review {
	return &Review{c}
}

// AddComments ...
func (r *Review) AddComments(ctx context.Context, input platform.Data) error {
	comments := make([]CommentData, len(input.RawDatas))
	for i, data := range input.RawDatas {
		var text string
		if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}
		comments[i] = CommentData{
			Body:      text,
			CommitID:  env.GithubCommitID,
			Path:      data.FilePath,
			StartLine: data.LineNum,
			Line:      data.LineNum + 1, // TODO: これで本当に良いか検討
		}
	}
	log.PrintJSON("[]CommentData", comments)

	var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := r.client.CreateComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if multiErr != nil {
		return multiErr
	}
	return nil
}
