//go:generate moq -rm -out $GOPATH/app/src/test/mock/bitbucket/$GOFILE -pkg mock . Client
package bitbucket

import (
	"context"
	"time"
)

// Client ...
type Client interface {
	GetComments(ctx context.Context) ([]Comment, error)
	PostComment(ctx context.Context, data CommentData) error
	UpsertReport(ctx context.Context, reportID string, data ReportData) error
	GetReport(ctx context.Context, reportID string) (*AnnotationResponse, error)
	DeleteReport(ctx context.Context, reportID string) error
	BulkUpsertAnnotations(ctx context.Context, datas []AnnotationData, reportID string) error
}

// request

// ---------------------

// NOTE: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
// CommentData ...
type CommentData struct {
	Content Content `json:"content,omitempty"`
	Inline  Inline  `json:"inline,omitempty"`
}

type Content struct {
	Type   string `json:"type,omitempty"`
	Raw    string `json:"raw,omitempty"`
	Markup string `json:"markup,omitempty"`
	HTML   string `json:"html,omitempty"`
}

type Inline struct {
	Path string `json:"path,omitempty"`
	From int    `json:"from,omitempty"`
	To   uint   `json:"to,omitempty"`
}

// ---------------------

// NOTE: https://developer.atlassian.com/cloud/bitbucket/rest/api-group-reports/#api-repositories-workspace-repo-slug-commit-commit-reports-reportid-put
type ReportData struct {
	UUID              string `json:"uuid,omitempty"`
	Title             string `json:"title,omitempty"`
	Details           string `json:"details,omitempty"`
	ExternalID        string `json:"external_id,omitempty"`
	Reporter          string `json:"reporter,omitempty"`
	Link              string `json:"link,omitempty"`
	RemoteLinkEnabled bool   `json:"remote_link_enabled,omitempty"`
	LogoURL           string `json:"logo_url,omitempty"`
	ReportType        string `json:"report_type,omitempty"` // SECURITY, COVERAGE, TEST, BUG
	Result            string `json:"result,omitempty"`      // PASSED, FAILED, PENDING
	Data              []Data `json:"data,omitempty"`
	CreatedOn         string `json:"created_on,omitempty"`
	UpdatedOn         string `json:"updated_on,omitempty"`
}

type Data struct {
	Type  string `json:"type"` // BOOLEAN, DATE, DURATION, LINK, NUMBER, PERCENTAGE, TEXT
	Title string `json:"title"`
	Value Value  `json:"value"`
}

type Value struct { // 好きなオブジェクトでOK
}

type AnnotationResponse struct {
	UUID string `json:"uuid,omitempty"`
}

type AnnotationData struct {
	UUID           string `json:"uuid,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	Type           string `json:"type,omitempty"`
	AnnotationType string `json:"annotation_type,omitempty"` // VULNERABILITY, CODE_SMELL, BUG
	Path           string `json:"path,omitempty"`
	Line           uint   `json:"line,omitempty"`
	Summary        string `json:"summary,omitempty"`
	Details        string `json:"details,omitempty"`
	Result         string `json:"result,omitempty"`   // PASSED, FAILED, IGNORED, SKIPPED
	Severity       string `json:"severity,omitempty"` // HIGH, MEDIUM, LOW, CRITICAL
	Link           string `json:"link,omitempty"`
	CreatedOn      string `json:"created_on,omitempty"`
	UpdatedOn      string `json:"updated_on,omitempty"`
}

// response
// ------------------

type PullRequestComments struct {
	Size     int       `json:"size"`
	Page     int       `json:"page"`
	Pagelen  int       `json:"pagelen"`
	Next     string    `json:"next,omitempty"`
	Previous string    `json:"previous"`
	Values   []Comment `json:"values"`
}

type Comment struct {
	ID          int       `json:"id"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	Content     Content   `json:"content"`
	User        User      `json:"user"`
	Deleted     bool      `json:"deleted"`
	Inline      Inline    `json:"inline"`
	Type        string    `json:"type"`
	PullRequest PR        `json:"pullrequest"`
	// Resolution  struct {
	// 	Type string `json:""`
	// }
	Pending bool `json:"pending"`
}

type User struct {
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
	// Links Link
	UUID      string `json:"uuid"`
	AccountID string `json:"account_id"`
	Nickname  string `json:"nickname"`
}

// type Link struct {

// }

type PR struct {
	Type  string `json:"type"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}
