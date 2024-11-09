package bitbucket

import (
	"context"
	"errors"
	"fmt"
	"slices"

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

	if len(input.Contents) > 0 {
		if err := r.addAnnotations(ctx, input, reportID); err != nil {
			return fmt.Errorf("failed to exec r.addAnnotations(): %w", err)
		}
		if err := r.addComments(ctx, input); err != nil {
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

	if len(input.Contents) == 0 {
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

func (r *Review) addAnnotations(ctx context.Context, input platform.Data, reportID string) error {
	if len(input.Contents) == 0 {
		return fmt.Errorf("there is no data to annotation")
	}

	annotations := make([]AnnotationData, len(input.Contents))
	for i, data := range input.Contents {
		annotations[i] = AnnotationData{
			ExternalID:     fmt.Sprintf("pr-commentator-%03d", i+1), // NOTE: bulk annotations で一度に作成できるのは MAX 100件まで
			Path:           data.FilePath,
			Line:           data.LineNum,
			Summary:        fmt.Sprintf("%s find problem", data.Linter),
			Details:        fmt.Sprintf("%s（%s）", data.Message, data.Linter),
			AnnotationType: "BUG",
			Result:         "FAILED",
			Severity:       "HIGH",
		}
	}

	if err := r.client.BulkUpsertAnnotations(ctx, annotations, reportID); err != nil {
		return err
	}
	return nil
}

func (r *Review) addComments(ctx context.Context, input platform.Data) error {
	if len(input.Contents) == 0 {
		return fmt.Errorf("there is no data to comment")
	}

	existingComments, err := r.client.GetComments(ctx)
	if err != nil {
		return fmt.Errorf("failed to exec r.getComments(): %w", err)
	}

	existingCommentIDs := make([]string, 0)
	for _, v := range existingComments {
		if !v.Deleted {
			existingCommentIDs = append(existingCommentIDs, fmt.Sprintf("%s:%d:%s", v.Inline.Path, v.Inline.To, v.Content.Raw))
		}
	}

	comments := make([]CommentData, 0)
	for _, content := range input.Contents {
		var text string
		if content.CustomCommentText != nil {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n%s", *content.CustomCommentText)
		} else {
			text = fmt.Sprintf("[*Automatic PR Comment*]  \n*・File:* %s（%d）  \n*・Linter:* %s  \n*・Details:* %s", content.FilePath, content.LineNum, content.Linter, content.Message) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
		}

		commentID := fmt.Sprintf("%s:%d:%s", content.FilePath, content.LineNum, text)
		if !slices.Contains(existingCommentIDs, commentID) { // NOTE: すでに同じファイルの同じ行に同じコメントがある場合はコメントしないように制御
			comments = append(comments, CommentData{
				Content: Content{
					Raw: text,
				},
				Inline: Inline{
					Path: content.FilePath,
					To:   content.LineNum,
				},
			})
		}
	}

	var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := r.client.PostComment(ctx, comment); err != nil {
			multiErr = errors.Join(multiErr, err)
		}
	}
	if multiErr != nil {
		return multiErr
	}
	return nil
}
