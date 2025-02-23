package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" log request")
	io.WriteString(w, "hello its me")
}

func myName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/name/")
	io.WriteString(w, fmt.Sprintf("my name is %s", name))
}

func main() {
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/name/{name}", myName)

	http.ListenAndServe(":3000", nil)
}
