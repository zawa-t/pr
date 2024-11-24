package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/zawa-t/pr/reporter/src/env"
	"github.com/zawa-t/pr/reporter/src/platform/github"
	"github.com/zawa-t/pr/reporter/src/platform/http"
	"github.com/zawa-t/pr/reporter/src/platform/http/url"
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
			Add(http.RequestHeader.Accept, "application/vnd.github+json").
			Add("X-GitHub-Api-Version", "2022-11-28").
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.Github.APIToken)),
	)

	res, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if res.StatusCode != 201 {
		slog.Error("Failed to post comment.", "req", fmt.Sprintf("URL: %s, Request body: %+v", parsedURL.String(), data), "res", fmt.Sprintf("%d: %s\n", res.StatusCode, string(res.Body)))
		return fmt.Errorf("failed to post comment")
	}

	return nil
}

// GetPRComments ...
func (c *Custom) GetPRComments(ctx context.Context) ([]github.GetPRCommentResponse, error) {
	parsedURL, err := url.New(prCommentURL)
	if err != nil {
		return nil, fmt.Errorf("failed to exec url.New(): %w", err)
	}

	req, err := http.NewRequest(http.Method.GET, parsedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(
		http.Header().
			Add(http.RequestHeader.Accept, "application/vnd.github+json").
			Add("X-GitHub-Api-Version", "2022-11-28").
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.Github.APIToken)),
	)

	httpRes, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if httpRes.StatusCode != 200 {
		slog.Error("Failed to get comment.", "req", fmt.Sprintf("URL: %s", parsedURL.String()), "res", fmt.Sprintf("%d: %s\n", httpRes.StatusCode, string(httpRes.Body)))
		return nil, fmt.Errorf("failed to get comments")
	}

	var res []github.GetPRCommentResponse
	if err := json.Unmarshal(httpRes.Body, &res); err != nil {
		return nil, fmt.Errorf("failed to exec json.Unmarshal(): %w", err)
	}

	return res, nil
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
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.Github.APIToken)),
	)

	res, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if res.StatusCode != 200 {
		slog.Error("Failed to post comment.", "req", fmt.Sprintf("URL: %s, Request body: %+v", parsedURL.String(), data), "res", fmt.Sprintf("%d: %s\n", res.StatusCode, string(res.Body)))
		return fmt.Errorf("failed to post comment")
	}

	return nil
}

// CreateCheckRun ...
func (c *Custom) CreateCheckRun(ctx context.Context, data github.POSTCheckRuns) error {
	parsedURL, err := url.New(checkRunURL)
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
			Add(http.RequestHeader.Authorization, fmt.Sprintf("Bearer %s", env.Github.APIToken)),
	)

	res, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if res.StatusCode != 201 {
		slog.Error("Failed to create check run.", "req", fmt.Sprintf("URL: %s, Request body: %+v", parsedURL.String(), data), "res", fmt.Sprintf("%d: %s\n", res.StatusCode, string(res.Body)))
		return fmt.Errorf("failed to post comment")
	}

	return nil
}
