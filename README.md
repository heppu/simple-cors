# Simple CORS
Dead simple CORS handler for Go.
Works with http.Handler compatible packages.

## Example
```go
package main

import (
	"io"
	"net/http"

	"github.com/heppu/simple-cors"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", cors.CORS()(mux))
}
```
