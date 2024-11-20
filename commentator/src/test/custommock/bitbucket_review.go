package custommock

import (
	"context"

	"github.com/zawa-t/pr/commentator/src/errors"
	"github.com/zawa-t/pr/commentator/src/platform/bitbucket"
	mock "github.com/zawa-t/pr/commentator/src/test/mock/bitbucket"
)

var DefaultBitbucketReview = &mock.ClientMock{
	BulkUpsertAnnotationsFunc: func(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error { return nil },
	DeleteReportFunc:          func(ctx context.Context, reportID string) error { return nil },
	GetCommentsFunc: func(ctx context.Context) ([]bitbucket.Comment, error) {
		return []bitbucket.Comment{}, nil
	},
	GetReportFunc: func(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
		return nil, errors.ErrNotFound
	},
	PostCommentFunc:  func(ctx context.Context, data bitbucket.CommentData) error { return nil },
	UpsertReportFunc: func(ctx context.Context, reportID string, data bitbucket.ReportData) error { return nil },
}
