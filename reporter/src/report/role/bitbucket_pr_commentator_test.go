package role

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zawa-t/pr/reporter/src/platform/bitbucket"
	"github.com/zawa-t/pr/reporter/src/report"
	"github.com/zawa-t/pr/reporter/src/test/custommock"
)

func Test_role_newCommentData(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBitbucketPRCommentator(custommock.DefaultBitbucketClientMock).newCommentData(context.Background(), tt.input)
			if assert.NoError(t, err) {
				assert.Equal(t, got, tt.expected)
			}
		})
	}
}
