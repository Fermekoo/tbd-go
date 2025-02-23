package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" log request")
	io.WriteString(w, "hello its me")
}

func main() {
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3000", nil)
}
