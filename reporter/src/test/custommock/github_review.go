package custommock

import (
	"context"

	"github.com/zawa-t/pr/reporter/src/platform/github"
	mock "github.com/zawa-t/pr/reporter/src/test/mock/github"
)

var DefaultGithubReview = &mock.ClientMock{
	CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
		return nil
	},
}
