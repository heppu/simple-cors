# Dead Simple CORS
CORS handling made as easy as it gets.

You can use this package when you want to allow all cross origin requests. It works with every http.Handler compatible router. Since the package is just one function I suggest that you just copy the code instead of importing it.

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
	http.ListenAndServe(":8000", cors.CORS(mux))
}
```

## How?
By setting following headers:
```
"Access-Control-Allow-Origin"	   : origin of request or * as a fallback
"Access-Control-Allow-Methods"	   : "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"
"Access-Control-Allow-Headers"	   : "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
"Access-Control-Expose-Headers"    : same as above
"Access-Control-Allow-Credentials" : true
```

## Licence
MIT licensed. See the LICENSE file for details.
