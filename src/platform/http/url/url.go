package url

import (
	"fmt"
	"log/slog"
	"net/url"
)

// URL ...
type URL struct {
	netURL *url.URL
}

// New ...
func New(rawURL string) (*URL, error) {
	v, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &URL{v}, nil
}

// String ...
func (u *URL) String() string {
	return u.netURL.String()
}

// JoinPathWithNoError ...
func JoinPathWithNoError(baseURL string, elem ...string) string {
	u, err := url.JoinPath(baseURL, elem...)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to join the provided path elements(%v) to BaseURL.", elem))
	}
	return u
}
