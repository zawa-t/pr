package role

import (
	"context"

	"github.com/zawa-t/pr/commentator/src/log"
	"github.com/zawa-t/pr/commentator/src/platform/github"
	"github.com/zawa-t/pr/commentator/src/review"
)

// githubPRChecker ...
type githubPRChecker struct {
	client github.Client
}

// NewGithubPRChecker ...
func NewGithubPRChecker(c github.Client) *githubPRChecker {
	return &githubPRChecker{c}
}

// Review ...
func (g *githubPRChecker) Review(ctx context.Context, input review.Data) error {
	comments := make([]github.Comment, len(input.Contents))
	for i, content := range input.Contents {
		comments[i] = github.Comment{
			Body:      content.Text,
			Path:      content.FilePath,
			StartLine: content.LineNum,
			Line:      content.LineNum + 1, // TODO: これで本当に良いか検討
		}
	}

	data := github.ReviewData{
		Body:     "yyyyy",
		Event:    "COMMENT",
		Comments: comments,
	}

	log.PrintJSON("ReviewData", data)

	if err := g.client.CreateReview(ctx, data); err != nil {
		return err
	}
	return nil
}
