package custommock

import (
	"context"

	"github.com/zawa-t/pr/reporter/platform/http"
	"github.com/zawa-t/pr/reporter/test/mock"
)

var Client = func(statusCode int, cookies http.Cookies, body []byte) *mock.ClientMock {
	return &mock.ClientMock{
		SendFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: statusCode,
				Headers:    http.Header(), // TODO: 必要に応じて修正
				Cookies:    cookies,
				Body:       body,
			}, nil
		},
	}
}
