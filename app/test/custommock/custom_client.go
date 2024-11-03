package custommock

import (
	"context"

	"github.com/zawa-t/pr-commentator/platform"
	"github.com/zawa-t/pr-commentator/platform/bitbucket"
	"github.com/zawa-t/pr-commentator/test/mock"
)

var DefaultCustomClient = &mock.CustomClientMock{
	BulkUpsertAnnotationsFunc: func(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error { return nil },
	DeleteReportFunc:          func(ctx context.Context, reportID string) error { return nil },
	GetReportFunc: func(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
		return nil, platform.ErrNotFound
	},
	PostCommentFunc:  func(ctx context.Context, data bitbucket.CommentData) error { return nil },
	UpsertReportFunc: func(ctx context.Context, reportID string, data bitbucket.ReportData) error { return nil },
}
