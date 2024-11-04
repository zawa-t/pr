package bitbucket

import (
	"context"
	"errors"
	"fmt"

	"github.com/zawa-t/pr-commentator/log"
	"github.com/zawa-t/pr-commentator/platform"
)

type pullRequest struct {
	review Review
}

func NewPullRequest(r Review) *pullRequest {
	return &pullRequest{r}
}

func (pr *pullRequest) AddComments(ctx context.Context, input platform.Data) error {
	reportID := fmt.Sprintf("pr-commentator-%s", input.Name)

	if err := pr.createReport(ctx, input, reportID); err != nil {
		return fmt.Errorf("failed to exec p.createReport(): %w", err)
	}

	if len(input.RawDatas) > 0 {
		if err := pr.addComments(ctx, input, reportID); err != nil {
			return fmt.Errorf("failed to exec p.addComments(): %w", err)
		}
	}

	return nil
}

func (pr *pullRequest) createReport(ctx context.Context, input platform.Data, reportID string) error {
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
	existingReport, err := pr.review.GetReport(ctx, reportID)
	if err != nil && err.Error() != platform.ErrNotFound.Error() { // TODO: err.Error() != platform.ErrNotFound.Error に errors.Is() が使えるやり方を検討
		return fmt.Errorf("failed to exec p.client.getReport(): %w", err)
	}
	if existingReport != nil {
		if err := pr.review.DeleteReport(ctx, reportID); err != nil {
			return fmt.Errorf("failed to exec p.client.deleteReport(): %w", err)
		}
	}

	if err := pr.review.UpsertReport(ctx, reportID, reportData); err != nil {
		return fmt.Errorf("failed to exec p.client.upsertReport(): %w", err)
	}
	return nil
}

func (pr *pullRequest) addComments(ctx context.Context, input platform.Data, reportID string) error {
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

		annotations[i] = AnnotationData{
			ExternalID:     fmt.Sprintf("pr-commentator-%03d", i+1), // NOTE: bulk annotations で一度に作成できるのは MAX 100件まで
			Path:           data.FilePath,
			Line:           data.LineNum,
			Summary:        fmt.Sprintf("%s（%s）", data.Summary, data.Linter),
			Details:        fmt.Sprintf("%s（%s）", data.Details, data.Linter),
			AnnotationType: "BUG",
			Result:         "FAILED",
			Severity:       "HIGH",
		}
	}

	log.PrintJSON("[]CommentData", comments)
	log.PrintJSON("[]AnnotationData", comments)

	var multiErr error // NOTE: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := pr.review.PostComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if err := pr.review.BulkUpsertAnnotations(ctx, annotations, reportID); err != nil {
		multiErr = errors.Join(multiErr, err)
	}

	if multiErr != nil {
		return multiErr
	}
	return nil
}
