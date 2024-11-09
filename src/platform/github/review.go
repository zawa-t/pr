package github

import (
	"context"
	"fmt"

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
	// comments := make([]CommentData, len(input.Contents))
	// for i, data := range input.Contents {
	// 	var text string
	// 	if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
	// 		text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
	// 	} else {
	// 		text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
	// 	}
	// 	comments[i] = CommentData{
	// 		Body:      text,
	// 		CommitID:  env.GithubCommitID,
	// 		Path:      data.FilePath,
	// 		StartLine: data.LineNum,
	// 		Line:      data.LineNum + 1, // TODO: これで本当に良いか検討
	// 	}
	// }
	// log.PrintJSON("[]CommentData", comments)

	// var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	// for _, comment := range comments {
	// 	if err := r.client.CreateComment(ctx, comment); err != nil {
	// 		multiErr = errors.Join(multiErr, err)
	// 	}
	// }
	// if multiErr != nil {
	// 	return multiErr
	// }
	// return nil
	return r.AddReviewComments(ctx, input)
}

// AddReviewComments ...
func (r *Review) AddReviewComments(ctx context.Context, input platform.Data) error {
	comments := make([]Comment, len(input.Contents))
	for i, data := range input.Contents {
		var text string
		if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}
		comments[i] = Comment{
			Body:      text,
			Path:      data.FilePath,
			StartLine: data.LineNum,
			Line:      data.LineNum + 1, // TODO: これで本当に良いか検討
		}
	}

	data := ReviewData{
		Body:     "xxxxxxxx",
		Event:    "REQUEST_CHANGES",
		Comments: comments,
	}

	log.PrintJSON("ReviewData", data)

	if err := r.client.CreateReview(ctx, data); err != nil {
		return err
	}
	return nil
}
