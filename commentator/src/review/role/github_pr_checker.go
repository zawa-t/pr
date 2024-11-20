package role

import (
	"context"
	"fmt"

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
	for i, data := range input.Contents {
		var text string
		if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}
		comments[i] = github.Comment{
			Body:      text,
			Path:      data.FilePath,
			StartLine: data.LineNum,
			Line:      data.LineNum + 1, // TODO: これで本当に良いか検討
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
