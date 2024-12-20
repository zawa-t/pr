package http

import (
	"bytes"
	"encoding/json"

	"github.com/zawa-t/pr/src/platform/http/url"
)

// Request ...
type Request struct {
	method method
	url    *url.URL
	body   *bytes.Buffer
	header header

	basicAuth *basicAuth
}

// NewRequest ...
func NewRequest(method method, url *url.URL, body any) (*Request, error) {
	b, err := newRequestBody(body)
	if err != nil {
		return nil, err
	}
	return &Request{method, url, b, nil, nil}, nil
}

func (r *Request) SetHeader(header header) {
	r.header = header
}

func (r *Request) SetBasicAuth(id, password string) {
	r.basicAuth = &basicAuth{id, password}
}

func (r *Request) hasHeader() bool {
	return len(r.header) > 0
}

type method string

var Method = struct {
	POST   method
	PUT    method
	DELETE method
	GET    method
}{
	POST:   "POST",
	PUT:    "PUT",
	DELETE: "DELETE",
	GET:    "GET",
}

func newRequestBody(body any) (*bytes.Buffer, error) {
	if body == nil {
		return bytes.NewBuffer(make([]byte, 0)), nil
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
