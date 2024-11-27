package role

import (
	"context"
	"testing"
	"time"

	"github.com/zawa-t/pr/src/platform/bitbucket"
	"github.com/zawa-t/pr/src/report"
	"github.com/zawa-t/pr/src/test/custommock"

	"github.com/stretchr/testify/assert"
)

func Test_role_bitbucketPRCommentator_newCommentData(t *testing.T) {
	type testCase struct {
		name     string
		input    report.Data
		expected []bitbucket.CommentData
	}

	tests := []testCase{
		{
			name: "正常系",
			input: report.Data{
				Name: "golangci-lint",
				Contents: []report.Content{
					{
						ID:        report.ReNewID("sample/main.go", 20, report.DefaultMessage("sample/main.go", 20, "golangci-lint", "sample text").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   20,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 20, "golangci-lint", "sample text"),
					},
					{
						ID:        report.ReNewID("sample/main.go", 11, report.DefaultMessage("sample/main.go", 11, "golangci-lint", "sample text1").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   11,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 11, "golangci-lint", "sample text1"),
					},
					{
						ID:        report.ReNewID("sample/main.go", 31, report.DefaultMessage("sample/main.go", 31, "golangci-lint", "sample text3").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   31,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 31, "golangci-lint", "sample text3"),
					},
				},
			},
			expected: []bitbucket.CommentData{
				{
					Content: bitbucket.Content{
						Raw: report.DefaultMessage("sample/main.go", 20, "golangci-lint", "sample text").String(),
					},
					Inline: bitbucket.Inline{
						Path: "sample/main.go",
						To:   20,
					},
				},
				{
					Content: bitbucket.Content{
						Raw: report.DefaultMessage("sample/main.go", 31, "golangci-lint", "sample text3").String(),
					},
					Inline: bitbucket.Inline{
						Path: "sample/main.go",
						To:   31,
					},
				},
			},
		},
	}

	bitbucketClientMock := custommock.DefaultBitbucketClientMock
	bitbucketClientMock.GetCommentsFunc = func(ctx context.Context) ([]bitbucket.Comment, error) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBitbucketPRCommentator(bitbucketClientMock).newCommentData(context.Background(), tt.input)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
