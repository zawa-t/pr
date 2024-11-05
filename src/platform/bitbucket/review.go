//go:generate moq -rm -out $GOPATH/app/src/test/mock/bitbucket/$GOFILE -pkg mock . Review
package bitbucket

import "context"

type Review interface {
	PostComment(ctx context.Context, data CommentData) error
	UpsertReport(ctx context.Context, reportID string, data ReportData) error
	GetReport(ctx context.Context, reportID string) (*AnnotationResponse, error)
	DeleteReport(ctx context.Context, reportID string) error
	BulkUpsertAnnotations(ctx context.Context, datas []AnnotationData, reportID string) error
}

// NOTE: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type CommentData struct {
	Content Content `json:"content,omitempty"`
	Inline  Inline  `json:"inline,omitempty"`
}

type Content struct {
	Raw    string `json:"raw,omitempty"`
	Markup string `json:"markup,omitempty"`
	HTML   string `json:"html,omitempty"`
}

type Inline struct {
	Path string `json:"path,omitempty"`
	To   uint   `json:"to,omitempty"`
}
