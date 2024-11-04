//go:generate moq -rm -out $GOPATH/app/test/mock/github/$GOFILE -pkg mock . Review
package github

import "context"

// Review ...
type Review interface {
	CreateComment(ctx context.Context, data CommentData) error
}

// CommentData ...
// 参考：https://docs.github.com/ja/rest/pulls/comments?apiVersion=2022-11-28#create-a-review-comment-for-a-pull-request
type CommentData struct {
	Body      string `json:"body"`       // example:"Great stuff!"
	CommitID  string `json:"commit_id"`  // example:"6dcb09b5b57875f334f61aebed695e2e4193db5e"
	Path      string `json:"path"`       // example:"file1.txt"
	StartLine uint   `json:"start_line"` // example:1
	StartSide string `json:"start_side"` // example:"RIGHT"
	Line      uint   `json:"line"`       // example:2
	Side      string `json:"side"`       // example:"RIGHT"
}
