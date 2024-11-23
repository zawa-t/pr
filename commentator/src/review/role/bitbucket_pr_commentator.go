package role

import (
	"context"
	stdErr "errors"
	"fmt"
	"slices"

	"github.com/zawa-t/pr/commentator/src/platform/bitbucket"
	"github.com/zawa-t/pr/commentator/src/review"
)

// bitbucketPRCommentator ...
type bitbucketPRCommentator struct {
	client bitbucket.Client
}

// NewBitbucketPRCommentator ...
func NewBitbucketPRCommentator(c bitbucket.Client) *bitbucketPRCommentator {
	return &bitbucketPRCommentator{c}
}

// Review ...
func (b *bitbucketPRCommentator) Review(ctx context.Context, input review.Data) error {
	reportID := fmt.Sprintf("pr-commentator-%s", input.Name)

	if err := b.createReport(ctx, input, reportID); err != nil {
		return fmt.Errorf("failed to exec r.createReport(): %w", err)
	}

	if len(input.Contents) > 0 {
		if err := b.addAnnotations(ctx, input, reportID); err != nil {
			return fmt.Errorf("failed to exec r.addAnnotations(): %w", err)
		}
		if err := b.addComments(ctx, input); err != nil {
			return fmt.Errorf("failed to exec r.addComments(): %w", err)
		}
	}

	return nil
}

func (b *bitbucketPRCommentator) createReport(ctx context.Context, input review.Data, reportID string) error {
	reportData := bitbucket.ReportData{
		Title:      fmt.Sprintf("[%s] PR-Commentator report", input.Name),
		Details:    "This report generated for you by pr-commentator.", // TODO: 内容については要検討
		ReportType: "TEST",
	}

	if len(input.Contents) == 0 {
		reportData.Result = "PASSED"
	} else {
		reportData.Result = "FAILED"
	}

	// TODO: 必要なさそうなので、一度コメントアウトして本当に不要であることが確認できたら削除する
	// // 存在確認
	// existingReport, err := b.client.GetReport(ctx, reportID)
	// if err != nil && err.Error() != errors.ErrNotFound.Error() { // TODO: err.Error() != platform.ErrNotFound.Error に errors.Is() が使えるやり方を検討
	// 	return fmt.Errorf("failed to exec r.client.getReport(): %w", err)
	// }
	// if existingReport != nil {
	// 	if err := b.client.DeleteReport(ctx, reportID); err != nil {
	// 		return fmt.Errorf("failed to exec r.client.deleteReport(): %w", err)
	// 	}
	// }

	if err := b.client.UpsertReport(ctx, reportID, reportData); err != nil {
		return fmt.Errorf("failed to exec r.client.upsertReport(): %w", err)
	}
	return nil
}

func (b *bitbucketPRCommentator) addAnnotations(ctx context.Context, input review.Data, reportID string) error {
	if len(input.Contents) == 0 {
		return fmt.Errorf("there is no data to annotation")
	}

	annotations := make([]bitbucket.AnnotationData, len(input.Contents))
	for i, data := range input.Contents {
		annotations[i] = bitbucket.AnnotationData{
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

	if err := b.client.BulkUpsertAnnotations(ctx, annotations, reportID); err != nil {
		return err
	}
	return nil
}

func (b *bitbucketPRCommentator) addComments(ctx context.Context, input review.Data) error {
	if len(input.Contents) == 0 {
		return fmt.Errorf("there is no data to comment")
	}

	existingComments, err := b.client.GetComments(ctx)
	if err != nil {
		return fmt.Errorf("failed to exec r.getComments(): %w", err)
	}

	existingCommentIDs := make([]review.ID, 0)
	for _, v := range existingComments {
		if !v.Deleted {
			existingCommentIDs = append(existingCommentIDs, review.ReNewID(v.Inline.Path, v.Inline.To, v.Content.Raw))
		}
	}

	comments := make([]bitbucket.CommentData, 0)
	for _, content := range input.Contents {
		if !slices.Contains(existingCommentIDs, content.ID) { // NOTE: すでに同じファイルの同じ行に同じコメントがある場合はコメントしないように制御
			comments = append(comments, bitbucket.CommentData{
				Content: bitbucket.Content{
					Raw: content.Message.String(),
				},
				Inline: bitbucket.Inline{
					Path: content.FilePath,
					To:   content.LineNum,
				},
			})
		}
	}

	var multiErr error // MEMO: 一部の処理が失敗しても残りの処理を進めたいため、エラーはすべての処理がおわってからハンドリング
	for _, comment := range comments {
		if err := b.client.PostComment(ctx, comment); err != nil {
			multiErr = stdErr.Join(multiErr, err)
		}
	}
	if multiErr != nil {
		return multiErr
	}
	return nil
}
