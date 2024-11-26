//go:generate moq -rm -out $GOPATH/src/reporter/test/mock/bitbucket/$GOFILE -pkg mock . Client
package bitbucket

import (
	"context"
)

// Client ...
type Client interface {
	GetComments(ctx context.Context) ([]Comment, error)
	PostComment(ctx context.Context, data CommentData) error
	UpsertReport(ctx context.Context, reportID string, data ReportData) error
	GetReport(ctx context.Context, reportID string) (*AnnotationResponse, error)
	DeleteReport(ctx context.Context, reportID string) error
	BulkUpsertAnnotations(ctx context.Context, datas []AnnotationData, reportID string) error
}
