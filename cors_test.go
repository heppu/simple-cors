package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "https://golang.org/", nil)
	req.Header.Set(origin, "localhost")
	if err != nil {
		panic(err)
	}
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wo := w.Header().Get(allow_origin)
		ro := r.Header.Get(origin)
		if wo != ro {
			t.Fatalf("Wrong Allow-Origin '%s' expected '%s'", wo, ro)
		}

		wm := w.Header().Get(allow_methods)
		if wm != methods {
			t.Fatalf("Wrong Allow-Methods '%s' expected '%s'", wm, methods)
		}

		wh := w.Header().Get(allow_headers)
		if wh != headers {
			t.Fatalf("Wrong Allow-Header '%s' expected '%s'", wh, headers)
		}
	})

	CORS(h).ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Fatal(http.StatusFound)
	}
}
