package http

// Response is a struct that represents a HTTP response.
type Response struct {
	StatusCode int
	Headers    header
	Cookies    Cookies
	Body       []byte
}
