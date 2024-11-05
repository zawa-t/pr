package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zawa-t/pr-commentator/src/platform/github"
	"github.com/zawa-t/pr-commentator/src/platform/http"
)

// Custom provides client for HTTP request.
type Custom struct {
	httpClient http.Client
}

// NewCustomClient creates a new client from options.
func NewCustomClient(hc http.Client) *Custom {
	return &Custom{hc}
}

// CreateComment ...
func (c *Custom) CreateComment(ctx context.Context, data github.CommentData) error {
	url, err := http.NewURL(prCommentURL)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.POST, url, data)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(http.Header().Add(http.RequestHeader.ContentType, http.ApplicationJSON))
	req.SetHeader(http.Header().Add(http.RequestHeader.Accept, "application/vnd.github+json"))
	req.SetHeader(http.Header().Add("X-GitHub-Api-Version", "2022-11-28"))
	req.SetBasicAuth("", "") // TODO: 何を設定すべきか検討
	/*
		(参考)
		-H "Accept: application/vnd.github+json" \
		-H "Authorization: Bearer <YOUR-TOKEN>" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
	*/

	res, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if res.StatusCode != 201 {
		slog.Error("Failed to post comment.", "req", data, "res", fmt.Sprintf("%d: %s\n", res.StatusCode, string(res.Body)))
		return fmt.Errorf("failed to post comment")
	}

	return nil
}
