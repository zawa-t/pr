package role

import (
	"context"
	"fmt"
	"slices"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/log"
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

	existingCommentIDs := make([]review.ID, 0)
	for _, v := range existingComments {
		existingCommentIDs = append(existingCommentIDs, review.ReNewID(v.Path, uint(v.StartLine), v.Body))
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

	log.PrintJSON("comments", comments)

	// var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	// for _, comment := range comments {
	// 	if err := g.client.CreateComment(ctx, comment); err != nil {
	// 		multiErr = errors.Join(multiErr, err)
	// 	}
	// }
	// if multiErr != nil {
	// 	return multiErr
	// }
	return nil
}
