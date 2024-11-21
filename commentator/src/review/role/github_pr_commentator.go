package role

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/platform/github"
	"github.com/zawa-t/pr/commentator/src/review"
)

// githubPRCommentator ...
type githubPRCommentator struct {
	client github.Client
}

// NewGithubPRCommentator ...
func NewGithubPRCommentator(c github.Client) *githubPRCommentator {
	return &githubPRCommentator{c}
}

// Review ...
func (g *githubPRCommentator) Review(ctx context.Context, input review.Data) error {
	if len(input.Contents) == 0 {
		return fmt.Errorf("there is no data to comment")
	}

	existingComments, err := g.client.GetPRComments(ctx)
	if err != nil {
		return fmt.Errorf("failed to exec r.GetPRComments(): %w", err)
	}

	existingCommentIDs := make([]string, 0)
	for _, v := range existingComments {
		existingCommentIDs = append(existingCommentIDs, fmt.Sprintf("%s:%d:%s", v.Path, v.StartLine, v.Body))
	}

	comments := make([]github.CommentData, len(input.Contents))
	for i, data := range input.Contents {
		var text string
		if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}

		commentID := fmt.Sprintf("%s:%d:%s", data.FilePath, data.LineNum, text)
		if !slices.Contains(existingCommentIDs, commentID) { // NOTE: すでに同じファイルの同じ行に同じコメントがある場合はコメントしないように制御
			comments[i] = github.CommentData{
				Body:        text,
				CommitID:    env.Github.CommitID,
				Path:        data.FilePath,
				StartLine:   data.LineNum,
				Line:        data.LineNum + 1, // TODO: これで本当に良いか検討
				Position:    5,
				SubjectType: "line",
			}
		}
	}

	var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := g.client.CreateComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if multiErr != nil {
		return multiErr
	}
	return nil
}
