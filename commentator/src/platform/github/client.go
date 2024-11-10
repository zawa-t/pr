//go:generate moq -rm -out $GOPATH/app/src/test/mock/github/$GOFILE -pkg mock . Client
package github

import "context"

// Client ...
type Client interface {
	CreateComment(ctx context.Context, data CommentData) error
	CreateReview(ctx context.Context, data ReviewData) error
}

// --------------------

// CommentData ...
// 参考：https://docs.github.com/ja/rest/pulls/comments?apiVersion=2022-11-28#create-a-review-comment-for-a-pull-request
type CommentData struct {
	Body        string  `json:"body"`                   // example:"Great stuff!"
	CommitID    string  `json:"commit_id"`              // example:"6dcb09b5b57875f334f61aebed695e2e4193db5e"
	Path        string  `json:"path"`                   // example:"file1.txt"
	StartLine   uint    `json:"start_line,omitempty"`   // example:1
	StartSide   *string `json:"start_side,omitempty"`   // example:"RIGHT" LEFT, RIGHT, side
	Line        uint    `json:"line,omitempty"`         // example:2  ※ Required unless using subject_type:file
	Side        *string `json:"side,omitempty"`         // example:"RIGHT" ※LEFT, RIGHT
	Position    uint    `json:"position,omitempty"`     // example:"RIGHT"
	InReplyTo   *uint   `json:"in_reply_to,omitempty"`  // example:"RIGHT"
	SubjectType string  `json:"subject_type,omitempty"` // example:"RIGHT" line, file
}

// --------------------

// ReviewData ...
// 参考：https://docs.github.com/ja/rest/pulls/reviews?apiVersion=2022-11-28#create-a-review-for-a-pull-request--code-samples
type ReviewData struct {
	CommitID string    `json:"commit_id,omitempty"` // example:"6dcb09b5b57875f334f61aebed695e2e4193db5e"
	Body     string    `json:"body,omitempty"`      // example:"Great stuff!"
	Event    string    `json:"event,omitempty"`     // example:"file1.txt", APPROVE, REQUEST_CHANGES, or COMMENT
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	Path      string `json:"path"`                 // example:"file1.txt"
	Position  uint   `json:"position,omitempty"`   // example:"RIGHT"
	Body      string `json:"body"`                 // example:"Great stuff!"
	Line      uint   `json:"line,omitempty"`       // example:2  ※ Required unless using subject_type:file
	Side      string `json:"side,omitempty"`       // example:"RIGHT" ※LEFT, RIGHT
	StartLine uint   `json:"start_line,omitempty"` // example:1
	StartSide string `json:"start_side,omitempty"` // example:"RIGHT" LEFT, RIGHT, side
}

// --------------------

// POSTCheckRuns ...
// 参考：https://docs.github.com/ja/rest/checks/runs?apiVersion=2022-11-28
type POSTCheckRuns struct {
	Name        string          `json:"name"`                   // example:
	HeadSHA     string          `json:"head_sha"`               // example:
	DetailsURL  string          `json:"details_url,omitempty"`  // example:
	ExternalID  string          `json:"external_id,omitempty"`  // example:
	Status      string          `json:"status,omitempty"`       // example:, queued, in_progress, completed, waiting, requested, pending
	StartedAt   string          `json:"started_at,omitempty"`   // example:
	Conclusion  string          `json:"conclusion,omitempty"`   // example:, action_required, cancelled, failure, neutral, success, skipped, stale, timed_out
	CompletedAt string          `json:"completed_at,omitempty"` // example:
	Output      CheckRunsOutput `json:"output,omitempty"`       // example:
	Actions     []Action        `json:"actions,omitempty"`      // example:
}

type CheckRunsOutput struct {
	Title       string       `json:"title"`                 // example:
	Summary     string       `json:"summary"`               // example:
	Text        string       `json:"text,omitempty"`        // example:
	Annotations []Annotation `json:"annotations,omitempty"` // example:
	Images      []Image      `json:"images,omitempty"`      // example:
}

type Annotation struct {
	Path            string `json:"path"`                   // example:
	StartLine       int    `json:"start_line"`             // example:
	EndLine         int    `json:"end_line"`               // example:
	StartColumn     int    `json:"start_column,omitempty"` // example:
	EndColumn       int    `json:"end_column,omitempty"`   // example:
	AnnotationLevel string `json:"annotation_level"`       // example:, notice, warning, failure
	Message         string `json:"message"`                // example:
	Title           string `json:"title,omitempty"`        // example:
	RawDetails      string `json:"raw_details,omitempty"`  // example:
}

type Image struct {
	Alt      string `json:"alt"`               // example:
	ImageURL string `json:"image_url"`         // example:
	Caption  string `json:"caption,omitempty"` // example:
}

type Action struct {
	Label       string `json:"label"`       // example:
	Description string `json:"description"` // example:
	Identifier  string `json:"identifier"`  // example:
}
