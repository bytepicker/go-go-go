// main
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Accessing with %q", html.EscapeString(r.Method+"/"+r.Proto))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
