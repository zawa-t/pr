package bitbucket

import (
	"context"
	"errors"
	"fmt"

	"github.com/zawa-t/pr-commentator/src/log"
	"github.com/zawa-t/pr-commentator/src/platform"
)

// Review ...
type Review struct {
	client Client
}

// NewReview ...
func NewReview(c Client) *Review {
	return &Review{c}
}

// AddComments ...
func (r *Review) AddComments(ctx context.Context, input platform.Data) error {
	reportID := fmt.Sprintf("pr-commentator-%s", input.Name)

	if err := r.createReport(ctx, input, reportID); err != nil {
		return fmt.Errorf("failed to exec r.createReport(): %w", err)
	}

	comments, err := r.getComments(ctx)
	if err != nil {
		return fmt.Errorf("failed to exec r.getComments(): %w", err)
	}

	log.PrintJSON("comments", comments)

	if len(input.RawDatas) > 0 {
		if err := r.addComments(ctx, input, reportID); err != nil {
			return fmt.Errorf("failed to exec r.addComments(): %w", err)
		}
	}

	return nil
}

func (r *Review) createReport(ctx context.Context, input platform.Data, reportID string) error {
	reportData := ReportData{
		Title:      fmt.Sprintf("[%s] PR-Commentator report", input.Name),
		Details:    "Meow-Meow! This report generated for you by pr-commentator.", // TODO: 内容については要検討
		ReportType: "TEST",
	}

	if len(input.RawDatas) == 0 {
		reportData.Result = "PASSED"
	} else {
		reportData.Result = "FAILED"
	}

	// 存在確認
	existingReport, err := r.client.GetReport(ctx, reportID)
	if err != nil && err.Error() != platform.ErrNotFound.Error() { // TODO: err.Error() != platform.ErrNotFound.Error に errors.Is() が使えるやり方を検討
		return fmt.Errorf("failed to exec r.client.getReport(): %w", err)
	}
	if existingReport != nil {
		if err := r.client.DeleteReport(ctx, reportID); err != nil {
			return fmt.Errorf("failed to exec r.client.deleteReport(): %w", err)
		}
	}

	if err := r.client.UpsertReport(ctx, reportID, reportData); err != nil {
		return fmt.Errorf("failed to exec r.client.upsertReport(): %w", err)
	}
	return nil
}

func (r *Review) getComments(ctx context.Context) (*PullRequestComments, error) {
	return r.client.GetComments(ctx)
}

func (r *Review) addComments(ctx context.Context, input platform.Data, reportID string) error {
	if len(input.RawDatas) == 0 {
		return fmt.Errorf("there is no data to comment")
	}

	comments := make([]CommentData, len(input.RawDatas))
	annotations := make([]AnnotationData, len(input.RawDatas))

	for i, data := range input.RawDatas {
		var text string
		if data.CustomCommentText != nil {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *data.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", data.FilePath, data.LineNum, data.Linter, data.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
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

		annotations[i] = AnnotationData{
			ExternalID:     fmt.Sprintf("pr-commentator-%03d", i+1), // NOTE: bulk annotations で一度に作成できるのは MAX 100件まで
			Path:           data.FilePath,
			Line:           data.LineNum,
			Summary:        fmt.Sprintf("%s（%s）", data.Message, data.Linter),
			Details:        fmt.Sprintf("%s（%s）", data.Message, data.Linter),
			AnnotationType: "BUG",
			Result:         "FAILED",
			Severity:       "HIGH",
		}
	}

	log.PrintJSON("[]CommentData", comments)
	log.PrintJSON("[]AnnotationData", comments)

	var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := r.client.PostComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if err := r.client.BulkUpsertAnnotations(ctx, annotations, reportID); err != nil {
		multiErr = errors.Join(multiErr, err)
	}

	if multiErr != nil {
		return multiErr
	}
	return nil
}
