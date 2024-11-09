package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zawa-t/pr-commentator/src/env"
	"github.com/zawa-t/pr-commentator/src/platform/github"
	"github.com/zawa-t/pr-commentator/src/platform/http"
	"github.com/zawa-t/pr-commentator/src/platform/http/url"
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
	parsedURL, err := url.New(prCommentURL)
	if err != nil {
		return fmt.Errorf("failed to exec url.New(): %w", err)
	}

	req, err := http.NewRequest(http.Method.POST, parsedURL, data)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(
		http.Header().
			Add(http.RequestHeader.ContentType, http.ApplicationJSON).
			Add(http.RequestHeader.Accept, "application/vnd.github+json").
			Add("X-GitHub-Api-Version", "2022-11-28").
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.GithubAPIToken)),
	)

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

// CreateReview ...
func (c *Custom) CreateReview(ctx context.Context, data github.ReviewData) error {
	parsedURL, err := url.New(prReviewURL)
	if err != nil {
		return fmt.Errorf("failed to exec url.New(): %w", err)
	}

	req, err := http.NewRequest(http.Method.POST, parsedURL, data)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(
		http.Header().
			Add(http.RequestHeader.Accept, "application/vnd.github+json").
			Add("X-GitHub-Api-Version", "2022-11-28").
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.GithubAPIToken)),
	)

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
