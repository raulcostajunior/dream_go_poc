package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h2>Hello, %q!</h2>", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "<p>platform: %q:%q", runtime.GOOS, runtime.GOARCH)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
