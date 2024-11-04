package custommock

import (
	"context"

	"github.com/zawa-t/pr-commentator/platform/github"
	mock "github.com/zawa-t/pr-commentator/test/mock/github"
)

var DefaultGithubReview = &mock.ReviewMock{
	CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
		return nil
	},
}
