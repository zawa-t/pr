package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/zawa-t/pr-commentator/env"
	"github.com/zawa-t/pr-commentator/platform"
	"github.com/zawa-t/pr-commentator/platform/bitbucket"
	"github.com/zawa-t/pr-commentator/platform/http"
)

// customClient provides client for HTTP request.
type customClient struct {
	httpClient http.Client
}

// NewCustomClient creates a new client from options.
func NewCustomClient(hc http.Client) *customClient {
	return &customClient{hc}
}

func (c *customClient) PostComment(ctx context.Context, data bitbucket.CommentData) error {
	url, err := http.NewURL(prCommentURL)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.POST, url, data)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(http.Header().Add(http.RequestHeader.ContentType, http.ApplicationJSON))
	req.SetBasicAuth(env.BitbucketUserName, env.BitbucketAppPassword)

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

func (c *customClient) UpsertReport(ctx context.Context, reportID string, data bitbucket.ReportData) error {
	url, err := http.NewURL(reportURL(reportID))
	if err != nil {
		return fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.PUT, url, data)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(http.Header().Add(http.RequestHeader.ContentType, http.ApplicationJSON))
	req.SetBasicAuth(env.BitbucketUserName, env.BitbucketAppPassword)

	httpRes, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if httpRes.StatusCode != 200 {
		slog.Error("Failed to put report.", "req", data, "res", fmt.Sprintf("%d: %s\n", httpRes.StatusCode, string(httpRes.Body)))
		return fmt.Errorf("failed to put report")
	}

	return nil
}

func (c *customClient) GetReport(ctx context.Context, reportID string) (*bitbucket.AnnotationResponse, error) {
	url, err := http.NewURL(reportURL(reportID))
	if err != nil {
		return nil, fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.GET, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetBasicAuth(env.BitbucketUserName, env.BitbucketAppPassword)

	httpRes, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if httpRes.StatusCode != 200 {
		slog.Warn("Failed to retrieve report.", "res", fmt.Sprintf("%d: %s\n", httpRes.StatusCode, string(httpRes.Body)))
		if httpRes.StatusCode == 404 {
			return nil, platform.ErrNotFound
		}
		return nil, fmt.Errorf("failed to retrieve report")
	}

	var res bitbucket.AnnotationResponse
	if err := json.Unmarshal(httpRes.Body, &res); err != nil {
		return nil, fmt.Errorf("failed to exec json.Unmarshal(): %w", err)
	}

	return &res, nil
}

func (c *customClient) DeleteReport(ctx context.Context, reportID string) error {
	url, err := http.NewURL(reportURL(reportID))
	if err != nil {
		return fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.DELETE, url, nil)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetBasicAuth(env.BitbucketUserName, env.BitbucketAppPassword)

	httpRes, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if httpRes.StatusCode != 200 && httpRes.StatusCode != 204 {
		slog.Error("Failed to delete report.", "res", fmt.Sprintf("%d: %s\n", httpRes.StatusCode, string(httpRes.Body)))
		return fmt.Errorf("failed to delete report")
	}

	return nil
}

func (c *customClient) BulkUpsertAnnotations(ctx context.Context, datas []bitbucket.AnnotationData, reportID string) error {
	url, err := http.NewURL(bulkAnnotationsURL(reportID))
	if err != nil {
		return fmt.Errorf("failed to exec http.NewURL(): %w", err)
	}

	req, err := http.NewRequest(http.Method.POST, url, datas)
	if err != nil {
		return fmt.Errorf("failed to exec http.NewRequest(): %w", err)
	}

	req.SetHeader(http.Header().Add(http.RequestHeader.ContentType, http.ApplicationJSON))
	req.SetBasicAuth(env.BitbucketUserName, env.BitbucketAppPassword)

	httpRes, err := c.httpClient.Send(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to exec c.httpClient.Send(): %w", err)
	}

	if httpRes.StatusCode != 200 {
		slog.Error("Failed to create annotation.", "res", datas, "res", fmt.Sprintf("%d: %s\n", httpRes.StatusCode, string(httpRes.Body)))
		return fmt.Errorf("failed to create annotation")
	}

	return nil
}
