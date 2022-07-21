package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
		fmt.Fprintf(w, "This web is written using golang")
	})

	http.ListenAndServe(":8080", nil)
}
