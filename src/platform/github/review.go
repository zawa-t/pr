package github

import (
	"context"
	"errors"

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
		comments[i] = CommentData{
			Body:      data.Details,
			CommitID:  env.GithubCommitID,
			Path:      data.FilePath,
			StartLine: data.LineNum - 1, // TODO: これで本当に良いか検討
			Line:      data.LineNum,
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
