//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Client
package http

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Client ...
type Client interface {
	Send(ctx context.Context, req *Request) (*Response, error)
}

type client struct {
	*http.Client
}

// NewClient creates a new client from options.
func NewClient(options ...Option) *client {
	// option := &Options{
	// 	Transport: &customTransport{
	// 		transport: http.DefaultTransport,
	// 	},
	// 	Timeout: 30 * time.Second,
	// }

	// for _, o := range options {
	// 	o(option)
	// }

	return &client{
		&http.Client{
			// Transport: option.Transport,
			// // Timeout:   option.Timeout,
		},
	}
}

func (c *client) Send(ctx context.Context, req *Request) (*Response, error) {
	httpReq, err := http.NewRequestWithContext(ctx, string(req.method), req.url.String(), req.body)
	if err != nil {
		return nil, err
	}

	if req.hasHeader() {
		httpReq.Header = http.Header(req.header)
	}

	if req.basicAuth != nil {
		httpReq.SetBasicAuth(req.basicAuth.id, req.basicAuth.password)
	}

	return c.send(httpReq)
}

func (c *client) send(req *http.Request) (*Response, error) {
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: res.StatusCode,
		Headers:    header(res.Header),
		Cookies:    res.Cookies(),
		Body:       b,
	}, nil
}

// Option is type of function for http client option.
type Option func(*Options)

// Options represents http client options.
type Options struct {
	Transport http.RoundTripper
	Timeout   time.Duration
}

// // 固定期間(Window)の構造体
// type Window struct {
// 	key           int64 // windowのキー
// 	requestCounts int   // window内のリクエスト数
// }

// // customTransport ...
// type customTransport struct {
// 	transport http.RoundTripper

// 	maxRetryCounts int // 最大リトライ数
// 	retryCounts    int // リトライした数

// 	maxRequestCounts int    // 単位時間あたりのリクエスト数上限
// 	perMilliSecond   int64  // 単位時間(ms)
// 	window           Window // 現在のwindow
// }

// func NewCustomTransport(wrt http.RoundTripper, maxRetryCounts int, maxRequestCounts int, perMilliSecond int64) *customTransport {
// 	return &customTransport{wrt, maxRetryCounts, 0, maxRequestCounts, perMilliSecond, Window{key: int64(0), requestCounts: 0}}
// }

// func (ct *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	start := time.Now()
// 	// Log request details
// 	slog.InfoContext(req.Context(), fmt.Sprintf("Start executing external API | %s %s", req.Method, req.URL.String()),
// 		"method", req.Method, "url", req.URL.String())

// 	for {
// 		now := time.Now().UnixMilli()
// 		cKey := now / ct.perMilliSecond

// 		// 新しい固定期間(window)になったらwaitなしでリクエストする
// 		if ct.window.key != cKey {
// 			ct.window = Window{
// 				key:           cKey,
// 				requestCounts: 0,
// 			}
// 			break
// 		}

// 		// 単位時間あたりのリクエスト数上限まではwaitなしでリクエストする
// 		if ct.window.requestCounts < ct.maxRequestCounts {
// 			break
// 		}

// 		// リクエスト数上限を超えていたら、waitする
// 		wait := ct.perMilliSecond - now%ct.perMilliSecond
// 		time.Sleep(time.Millisecond * time.Duration(wait))
// 	}

// 	ct.window.requestCounts++

// 	var resp *http.Response
// 	var err error
// 	for {
// 		resp, err = ct.transport.RoundTrip(req)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if resp != nil && resp.StatusCode < http.StatusInternalServerError {
// 			break
// 		}

// 		// リトライ数の上限チェック
// 		ct.retryCounts++
// 		if ct.retryCounts > ct.maxRetryCounts {
// 			break
// 		}

// 		// Exponential BackOff でウェイトを入れる
// 		time.Sleep(time.Second * time.Duration(math.Pow(2, float64(ct.retryCounts))))
// 	}
// 	ct.retryCounts = 0

// 	// Log response details
// 	latency := time.Since(start)
// 	if err == nil {
// 		slog.InfoContext(req.Context(), fmt.Sprintf("End executing external API | %s %s", req.Method, req.URL.String()),
// 			"method", req.Method, "url", req.URL.String(), "status", resp.Status, "latency", latency.String())
// 	} else {
// 		slog.WarnContext(req.Context(), fmt.Sprintf("End executing external API | %s %s", req.Method, req.URL.String()),
// 			"method", req.Method, "url", req.URL.String(), "latency", latency.String(), "error", err.Error())
// 	}
// 	return resp, err
// }
