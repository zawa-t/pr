package custommock

import (
	"context"

	"github.com/zawa-t/pr/reporter/platform/github"

	mock "github.com/zawa-t/pr/reporter/test/mock/github"
)

var DefaultGithubClientMock = &mock.ClientMock{
	CreateCheckRunFunc: func(ctx context.Context, data github.POSTCheckRuns) error { return nil },
	CreateCommentFunc:  func(ctx context.Context, data github.CommentData) error { return nil },
	CreateReviewFunc:   func(ctx context.Context, data github.ReviewData) error { return nil },
	GetPRCommentsFunc: func(ctx context.Context) ([]github.GetPRCommentResponse, error) {
		return []github.GetPRCommentResponse{}, nil
	},
}
