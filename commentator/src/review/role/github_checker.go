package role

import (
	"context"
	"fmt"

	"github.com/zawa-t/pr/commentator/src/env"
	"github.com/zawa-t/pr/commentator/src/log"
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
		Name:       "xxxx",
		HeadSHA:    env.Github.CommitID,
		Status:     "completed",
		Conclusion: "failure",
		Output: github.CheckRunsOutput{
			Title:   input.Name,
			Summary: "",
			Text:    "",
		},
	}

	annotations := make([]github.Annotation, len(input.Contents))
	for i, data := range input.Contents {
		var text string
		if data.CustomCommentText != nil { // HACK: bitbucketと同じ内容のため共通化したい
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}
		annotations[i] = github.Annotation{
			Path:            data.FilePath,
			StartLine:       int(data.LineNum),
			EndLine:         int(data.LineNum) + 1,
			AnnotationLevel: "warning",
			Message:         text,
			Title:           input.Name,
			// RawDetails:      "",
		}
	}
	postheckRuns.Output.Annotations = annotations

	log.PrintJSON("postheckRuns", postheckRuns)

	if err := g.client.CreateCheckRun(ctx, postheckRuns); err != nil {
		return err
	}
	return nil
}
