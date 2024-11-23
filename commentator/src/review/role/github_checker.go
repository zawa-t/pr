package role

import (
	"context"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/platform/github"
	"github.com/zawa-t/pr/commentator/src/review"
)

// githubChecker ...
type githubChecker struct {
	client github.Client
}

// NewGithubChecker ...
func NewGithubChecker(c github.Client) *githubChecker {
	return &githubChecker{c}
}

// Review ...
func (g *githubChecker) Review(ctx context.Context, input review.Data) error {
	postheckRuns := github.POSTCheckRuns{
		Name:       input.Name,
		HeadSHA:    env.Github.CommitID,
		Status:     "completed",
		Conclusion: "failure",
		Output: github.CheckRunsOutput{
			Title:   "Title",
			Summary: "Summary",
			Text:    "xxxxxxxxxxx",
		},
	}

	annotations := make([]github.Annotation, len(input.Contents))
	for i, content := range input.Contents {
		annotations[i] = github.Annotation{
			Path:            content.FilePath,
			StartLine:       int(content.LineNum),
			EndLine:         int(content.LineNum) + 1,
			AnnotationLevel: "warning",
			Message:         content.Message.String(),
			Title:           content.Linter,
			// RawDetails:      "",
		}
	}
	postheckRuns.Output.Annotations = annotations

	if err := g.client.CreateCheckRun(ctx, postheckRuns); err != nil {
		return err
	}
	return nil
}
