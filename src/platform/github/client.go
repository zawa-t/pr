//go:generate moq -rm -out $GOPATH/src/reporter/test/mock/github/$GOFILE -pkg mock . Client
package github

import "context"

// Client ...
type Client interface {
	CreateComment(ctx context.Context, data CommentData) error
	GetPRComments(ctx context.Context) ([]GetPRCommentResponse, error)
	CreateReview(ctx context.Context, data ReviewData) error
	CreateCheckRun(ctx context.Context, data POSTCheckRuns) error
}
