package role

import (
	"context"

	"github.com/zawa-t/pr/reporter/src/log"
	"github.com/zawa-t/pr/reporter/src/report"
)

// localCommentator ...
type localCommentator struct {
}

// NewLocalCommentator ...
func NewLocalCommentator() *localCommentator {
	return &localCommentator{}
}

// Report ...
func (pr *localCommentator) Report(ctx context.Context, input report.Data) error {
	log.PrintJSON("The following are the assumptions that will be submitted as report data.", input)
	return nil
}
