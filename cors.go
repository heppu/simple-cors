package cors

import "net/http"

const (
	options           string = "OPTIONS"
	allow_origin      string = "Access-Control-Allow-Origin"
	allow_methods     string = "Access-Control-Allow-Methods"
	allow_headers     string = "Access-Control-Allow-Headers"
	allow_credentials string = "Access-Control-Allow-Credentials"
	expose_headers    string = "Access-Control-Expose-Headers"
	credentials       string = "true"
	origin            string = "Origin"
	methods           string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"

	// If you want to expose some other headers add it here
	headers string = "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
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
	// Set allow origin to match origin of our request or fall back to *
	if o := r.Header.Get(origin); o != "" {
		w.Header().Set(allow_origin, o)
	} else {
		w.Header().Set(allow_origin, "*")
	}

	// Set other headers
	w.Header().Set(allow_headers, headers)
	w.Header().Set(allow_methods, methods)
	w.Header().Set(allow_credentials, credentials)
	w.Header().Set(expose_headers, headers)

	// If this was preflight options request let's write empty ok response and return
	if r.Method == options {
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}

	c.h.ServeHTTP(w, r)
}
