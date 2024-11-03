//go:generate moq -rm -out $GOPATH/app/test/mock/$GOFILE -pkg mock . CustomClient
package bitbucket

import (
	"context"
	"errors"
	"fmt"

	"github.com/zawa-t/pr-commentator/platform"
)

type CustomClient interface {
	PostComment(ctx context.Context, data CommentData) error
	UpsertReport(ctx context.Context, reportID string, data ReportData) error
	GetReport(ctx context.Context, reportID string) (*AnnotationResponse, error)
	DeleteReport(ctx context.Context, reportID string) error
	BulkUpsertAnnotations(ctx context.Context, datas []AnnotationData, reportID string) error
}

type pullRequest struct {
	client CustomClient
}

func NewPullRequest(cc CustomClient) *pullRequest {
	return &pullRequest{cc}
}

func (p *pullRequest) AddComments(ctx context.Context, input platform.Input) error {
	reportID := fmt.Sprintf("reviewcat-%s", input.Name)

	if err := p.createReport(ctx, input, reportID); err != nil {
		return fmt.Errorf("failed to exec p.createReport(): %w", err)
	}

	if len(input.Datas) > 0 {
		if err := p.addComments(ctx, input, reportID); err != nil {
			return fmt.Errorf("failed to exec p.addComments(): %w", err)
		}
	}

	return nil
}

func (p *pullRequest) createReport(ctx context.Context, input platform.Input, reportID string) error {
	reportData := ReportData{
		Title:      fmt.Sprintf("[%s] reviewcat report", input.Name),
		Details:    "Meow-Meow! This report generated for you by reviewcat.", // TODO: 内容については要検討
		ReportType: "TEST",
	}

	if len(input.Datas) == 0 {
		reportData.Result = "PASSED"
	} else {
		reportData.Result = "FAILED"
	}

	// 存在確認
	existingReport, err := p.client.GetReport(ctx, reportID)
	if err != nil && err.Error() != platform.ErrNotFound.Error() { // TODO: err.Error() != platform.ErrNotFound.Error に errors.Is() が使えるやり方を検討
		return fmt.Errorf("failed to exec p.client.getReport(): %w", err)
	}
	if existingReport != nil {
		if err := p.client.DeleteReport(ctx, reportID); err != nil {
			return fmt.Errorf("failed to exec p.client.deleteReport(): %w", err)
		}
	}

	if err := p.client.UpsertReport(ctx, reportID, reportData); err != nil {
		return fmt.Errorf("failed to exec p.client.upsertReport(): %w", err)
	}
	return nil
}

func (p *pullRequest) addComments(ctx context.Context, input platform.Input, reportID string) error {
	if len(input.Datas) == 0 {
		return fmt.Errorf("there is no data to comment")
	}

	comments := make([]CommentData, len(input.Datas))
	annotations := make([]AnnotationData, len(input.Datas))

	for i, data := range input.Datas {
		var text string
		if data.CustomCommentText != nil {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Details) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}

		comments[i] = CommentData{
			Content: Content{
				Raw: text,
			},
			Inline: Inline{
				Path: data.FilePath,
				To:   data.LineNum,
			},
		}
		fmt.Printf("comment: %+v\n", comments[i])

		annotations[i] = AnnotationData{
			ExternalID:     fmt.Sprintf("reviewcat-%03d", i+1), // NOTE: bulk annotations で一度に作成できるのは MAX 100件まで
			Path:           data.FilePath,
			Line:           data.LineNum,
			Summary:        fmt.Sprintf("%s（%s）", data.Summary, data.Linter),
			Details:        fmt.Sprintf("%s（%s）", data.Details, data.Linter),
			AnnotationType: "BUG",
			Result:         "FAILED",
			Severity:       "HIGH",
		}
	}

	var multiErr error // NOTE: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := p.client.PostComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if err := p.client.BulkUpsertAnnotations(ctx, annotations, reportID); err != nil {
		multiErr = errors.Join(multiErr, err)
	}

	if multiErr != nil {
		return multiErr
	}
	return nil
}
