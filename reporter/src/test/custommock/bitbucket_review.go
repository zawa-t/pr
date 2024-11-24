package custommock

import (
	"context"
	"time"

	"github.com/zawa-t/pr/reporter/src/errors"
	"github.com/zawa-t/pr/reporter/src/platform/bitbucket"
	"github.com/zawa-t/pr/reporter/src/report"
	mock "github.com/zawa-t/pr/reporter/src/test/mock/bitbucket"
)

var DefaultBitbucketClientMock = &mock.ClientMock{
	BulkUpsertAnnotationsFunc: func(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error { return nil },
	DeleteReportFunc:          func(ctx context.Context, reportID string) error { return nil },
	GetCommentsFunc: func(ctx context.Context) ([]bitbucket.Comment, error) {
		now := time.Now()
		comments := []bitbucket.Comment{
			{
				ID:        1,
				CreatedOn: now,
				UpdatedOn: now,
				Content: bitbucket.Content{
					Type:   "",
					Raw:    report.DefaultMessage("sample/main.go", 11, "golangci-lint", "sample text1").String(),
					Markup: "",
					HTML:   "",
				},
				User: struct {
					Type        string "json:\"type\""
					DisplayName string "json:\"display_name\""
					UUID        string "json:\"uuid\""
					AccountID   string "json:\"account_id\""
					Nickname    string "json:\"nickname\""
				}{
					Type:        "",
					DisplayName: "",
					UUID:        "",
					AccountID:   "",
					Nickname:    "",
				},
				Deleted: false,
				Inline: bitbucket.Inline{
					Path: "sample/main.go",
					From: 10,
					To:   11,
				},
				Type: "",
				PullRequest: struct {
					Type  string "json:\"type\""
					ID    int    "json:\"id\""
					Title string "json:\"title\""
				}{
					Type:  "",
					ID:    1,
					Title: "",
				},
				Pending: false,
			},
			{
				ID:        2,
				CreatedOn: now,
				UpdatedOn: now,
				Content: bitbucket.Content{
					Type:   "",
					Raw:    report.DefaultMessage("sample/main.go", 16, "golangci-lint", "sample text2").String(),
					Markup: "",
					HTML:   "",
				},
				User: struct {
					Type        string "json:\"type\""
					DisplayName string "json:\"display_name\""
					UUID        string "json:\"uuid\""
					AccountID   string "json:\"account_id\""
					Nickname    string "json:\"nickname\""
				}{
					Type:        "",
					DisplayName: "",
					UUID:        "",
					AccountID:   "",
					Nickname:    "",
				},
				Deleted: false,
				Inline: bitbucket.Inline{
					Path: "sample/main.go",
					From: 15,
					To:   16,
				},
				Type: "",
				PullRequest: struct {
					Type  string "json:\"type\""
					ID    int    "json:\"id\""
					Title string "json:\"title\""
				}{
					Type:  "",
					ID:    1,
					Title: "",
				},
				Pending: false,
			},
			{
				ID:        3,
				CreatedOn: now,
				UpdatedOn: now,
				Content: bitbucket.Content{
					Type:   "",
					Raw:    report.DefaultMessage("sample/main.go", 31, "golangci-lint", "sample text3").String(),
					Markup: "",
					HTML:   "",
				},
				User: struct {
					Type        string "json:\"type\""
					DisplayName string "json:\"display_name\""
					UUID        string "json:\"uuid\""
					AccountID   string "json:\"account_id\""
					Nickname    string "json:\"nickname\""
				}{
					Type:        "",
					DisplayName: "",
					UUID:        "",
					AccountID:   "",
					Nickname:    "",
				},
				Deleted: true,
				Inline: bitbucket.Inline{
					Path: "sample/main.go",
					From: 30,
					To:   31,
				},
				Type: "",
				PullRequest: struct {
					Type  string "json:\"type\""
					ID    int    "json:\"id\""
					Title string "json:\"title\""
				}{
					Type:  "",
					ID:    1,
					Title: "",
				},
				Pending: false,
			},
		}
		return comments, nil
	},
	GetReportFunc: func(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
		return nil, errors.ErrNotFound
	},
	PostCommentFunc:  func(ctx context.Context, data bitbucket.CommentData) error { return nil },
	UpsertReportFunc: func(ctx context.Context, reportID string, data bitbucket.ReportData) error { return nil },
}
