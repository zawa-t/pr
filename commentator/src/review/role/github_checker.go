package role

import (
	"context"
	"fmt"

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
			Title:   fmt.Sprintf("[%s] report some issues", input.Name),
			Summary: fmt.Sprintf("The total number of issues reported is %d.", len(input.Contents)),
			// Text: fmt.Sprintf("The total number of issues reported is %d.", len(input.Contents)),
		},
	}

	annotations := make([]github.Annotation, len(input.Contents))
	for i, content := range input.Contents {
		annotations[i] = github.Annotation{
			AnnotationLevel: "warning",
			Path:            content.FilePath,
			StartLine:       int(content.LineNum),
			EndLine:         int(content.LineNum),
			Title:           fmt.Sprintf("reported by [%s]", content.Linter),
			Message:         content.Message.String(),
			// RawDetails:      "",
		}
	}
	postheckRuns.Output.Annotations = annotations

	if err := g.client.CreateCheckRun(ctx, postheckRuns); err != nil {
		return err
	}
	return nil
}
