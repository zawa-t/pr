package role

import (
	"context"

	"github.com/zawa-t/pr/reporter/src/log"
	"github.com/zawa-t/pr/reporter/src/platform/github"
	"github.com/zawa-t/pr/reporter/src/report"
)

// githubPRReviewer ...
type githubPRReviewer struct {
	client github.Client
}

// NewGithubPRReviewer ...
func NewGithubPRReviewer(c github.Client) *githubPRReviewer {
	return &githubPRReviewer{c}
}

// Report ...
func (g *githubPRReviewer) Report(ctx context.Context, input report.Data) error {
	comments := make([]github.Comment, len(input.Contents))
	for i, content := range input.Contents {
		comments[i] = github.Comment{
			Body:      content.Message.String(),
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
