package custommock

import (
	"context"

	"github.com/zawa-t/pr-commentator/src/platform/github"
	mock "github.com/zawa-t/pr-commentator/src/test/mock/github"
)

var DefaultGithubReview = &mock.ReviewMock{
	CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
		return nil
	},
}
