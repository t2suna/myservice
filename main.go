package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Welcome My Page!\n")
	}

	http.HandleFunc("/new/", handler1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
