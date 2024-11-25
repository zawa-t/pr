package role

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zawa-t/pr/reporter/src/env"
	"github.com/zawa-t/pr/reporter/src/platform/github"
	"github.com/zawa-t/pr/reporter/src/report"
	"github.com/zawa-t/pr/reporter/src/test/custommock"
)

func Test_role_githubPRCommentator_newCommentData(t *testing.T) {
	type testCase struct {
		name     string
		input    report.Data
		expected []github.CommentData
	}

	tests := []testCase{
		{
			name: "正常系",
			input: report.Data{
				Name: "golangci-lint",
				Contents: []report.Content{
					{
						ID:        report.ReNewID("sample/main.go", 1, report.DefaultMessage("sample/main.go", 1, "golangci-lint", "sample text1").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   1,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 1, "golangci-lint", "sample text1"),
					},
					{
						ID:        report.ReNewID("sample/main.go", 2, report.DefaultMessage("sample/main.go", 2, "golangci-lint", "sample text2").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   2,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 2, "golangci-lint", "sample text2"),
					},
					{
						ID:        report.ReNewID("sample/main.go", 30, report.DefaultMessage("sample/main.go", 30, "golangci-lint", "sample text3").String()),
						Linter:    "golangci-lint",
						FilePath:  "sample/main.go",
						LineNum:   30,
						ColumnNum: 1,
						CodeLine:  "",
						Indicator: "",
						Message:   report.DefaultMessage("sample/main.go", 30, "golangci-lint", "sample text3"),
					},
				},
			},
			expected: []github.CommentData{
				{
					Body:        report.DefaultMessage("sample/main.go", 30, "golangci-lint", "sample text3").String(),
					CommitID:    env.Github.CommitID,
					Path:        "sample/main.go",
					StartLine:   30,
					Line:        31,
					Position:    5,
					SubjectType: "line",
				},
			},
		},
	}

	githubClientMock := custommock.DefaultGithubClientMock
	githubClientMock.GetPRCommentsFunc = func(ctx context.Context) ([]github.GetPRCommentResponse, error) {
		res := []github.GetPRCommentResponse{
			{
				Path:      "sample/main.go",
				Body:      report.DefaultMessage("sample/main.go", 1, "golangci-lint", "sample text1").String(),
				StartLine: 1,
			},
			{
				Path:      "sample/main.go",
				Body:      report.DefaultMessage("sample/main.go", 2, "golangci-lint", "sample text2").String(),
				StartLine: 2,
			},
		}
		return res, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGithubPRCommentator(githubClientMock).newCommentData(context.Background(), tt.input)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
