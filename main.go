package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
