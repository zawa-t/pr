package http

import "net/http"

type header http.Header

// Headers creates a new Headers.
func Header() header {
	return make(header)
}

func (h header) Add(key string, values ...string) header {
	h[key] = append(h[key], values...)
	return h
}

// // header is a Header of a HTTP request/response.
// type header struct {
// 	header  http.Header
// 	cookies Cookies
// }

// // NewHeader creates a new Header.
// func Header(header http.Header) header {
// 	return header{header: header, cookies: nil}
// }

var RequestHeader = struct {
	Accept        string
	ContentType   string
	Authorization string
}{
	Accept:        "Accept",
	ContentType:   "Content-Type",
	Authorization: "Authorization",
}

var (
	ApplicationJSON = "application/json"
)

// Cookies is the http cookies.
type Cookies []*http.Cookie

type basicAuth struct {
	id       string
	password string
}
