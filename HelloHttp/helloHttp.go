package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Bassel From Go Server. Heykal is the Host and the only host.: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
