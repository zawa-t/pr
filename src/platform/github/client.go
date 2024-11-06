//go:generate moq -rm -out $GOPATH/app/src/test/mock/github/$GOFILE -pkg mock . Client
package github

import "context"

// Client ...
type Client interface {
	CreateComment(ctx context.Context, data CommentData) error
}

// --------------------

// CommentData ...
// 参考：https://docs.github.com/ja/rest/pulls/comments?apiVersion=2022-11-28#create-a-review-comment-for-a-pull-request
type CommentData struct {
	Body        string `json:"body"`                   // example:"Great stuff!"
	CommitID    string `json:"commit_id"`              // example:"6dcb09b5b57875f334f61aebed695e2e4193db5e"
	Path        string `json:"path"`                   // example:"file1.txt"
	StartLine   uint   `json:"start_line,omitempty"`   // example:1
	StartSide   string `json:"start_side,omitempty"`   // example:"RIGHT" LEFT, RIGHT, side
	Line        uint   `json:"line,omitempty"`         // example:2  ※ Required unless using subject_type:file
	Side        string `json:"side,omitempty"`         // example:"RIGHT" ※LEFT, RIGHT
	Position    uint   `json:"position,omitempty"`     // example:"RIGHT"
	InReplyTo   uint   `json:"in_reply_to,omitempty"`  // example:"RIGHT"
	SubjectType string `json:"subject_type,omitempty"` // example:"RIGHT" line, file
}
