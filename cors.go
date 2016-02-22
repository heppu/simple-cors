package cors

import "net/http"

const (
	allow_origin  string = "Access-Control-Allow-Origin"
	allow_methods string = "Access-Control-Allow-Methods"
	allow_headers string = "Access-Control-Allow-Headers"
	origin        string = "Origin"
	methods       string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"
	headers       string = "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
)

type corsHandler struct {
	h http.Handler
}

func CORS() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &corsHandler{h}
	}
}

func (c *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if o := r.Header.Get(origin); o != "" {
		w.Header().Set(allow_origin, o)
		w.Header().Set(allow_methods, methods)
		w.Header().Set(allow_headers, headers)
	}
	c.h.ServeHTTP(w, r)
}
