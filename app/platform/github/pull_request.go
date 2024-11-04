package github

import (
	"context"

	"github.com/zawa-t/pr-commentator/log"
	"github.com/zawa-t/pr-commentator/platform"
)

// PullRequest ...
type PullRequest struct {
	review Review
}

// NewPullRequest ...
func NewPullRequest(r Review) *PullRequest {
	return &PullRequest{r}
}

// AddComments ...
func (pr *PullRequest) AddComments(ctx context.Context, input platform.Data) error {
	for _, data := range input.RawDatas {
		comment := CommentData{
			Body:      "",
			CommitID:  "",
			Path:      data.FilePath,
			StartLine: 2,
			StartSide: "RIGHT",
			Line:      data.LineNum,
			Side:      "RIGHT",
		}

		log.PrintJSON("CommentData", comment)
		if err := pr.review.CreateComment(ctx, comment); err != nil {
			return err
		}
	}
	return nil
}
