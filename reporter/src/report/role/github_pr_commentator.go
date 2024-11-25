package role

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/zawa-t/pr/reporter/src/env"
	"github.com/zawa-t/pr/reporter/src/platform/github"
	"github.com/zawa-t/pr/reporter/src/report"
)

// githubPRCommentator ...
type githubPRCommentator struct {
	client github.Client
}

// NewGithubPRCommentator ...
func NewGithubPRCommentator(c github.Client) *githubPRCommentator {
	return &githubPRCommentator{c}
}

// Report ...
func (g *githubPRCommentator) Report(ctx context.Context, input report.Data) error {
	comments, err := g.newCommentData(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to newCommentData(): %w", err)
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

func (g *githubPRCommentator) newCommentData(ctx context.Context, input report.Data) ([]github.CommentData, error) {
	if len(input.Contents) == 0 {
		return nil, fmt.Errorf("there is no data to comment")
	}

	existingComments, err := g.client.GetPRComments(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to GetPRComments(): %w", err)
	}

	existingCommentIDs := make([]report.ID, 0)
	for _, v := range existingComments {
		existingCommentIDs = append(existingCommentIDs, report.ReNewID(v.Path, uint(v.StartLine), v.Body))
	}

	comments := make([]github.CommentData, 0)
	for _, content := range input.Contents {
		if !slices.Contains(existingCommentIDs, content.ID) { // NOTE: すでに同じファイルの同じ行に同じコメントがある場合はコメントしないように制御
			comments = append(comments, github.CommentData{
				Body:        content.Message.String(),
				CommitID:    env.Github.CommitID,
				Path:        content.FilePath,
				StartLine:   content.LineNum,
				Line:        content.LineNum + 1, // TODO: これで本当に良いか検討
				Position:    5,
				SubjectType: "line",
			})
		}
	}
	return comments, nil
}
