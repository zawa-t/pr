package role

import (
	"context"

	"github.com/zawa-t/pr/commentator/src/log"
	"github.com/zawa-t/pr/commentator/src/review"
)

// localCommentator ...
type localCommentator struct {
}

// NewLocalCommentator ...
func NewLocalCommentator() *localCommentator {
	return &localCommentator{}
}

// Review ...
func (pr *localCommentator) Review(ctx context.Context, input review.Data) error {
	log.PrintJSON("The following are the assumptions that will be submitted as review data.", input)
	return nil
}
