package main

import (
	"fmt"
	"io"
	"net/http"
)

func dontPanic(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "42" {
		io.WriteString(w, `{"42": "a resposta para a vida, o universo e tudo mais"}`)
		fmt.Println(r.Method,
			"Authorization:",
			r.Header.Get("Authorization"),
			r.URL)
		return
	}
	fmt.Println(r.Method, r.URL)
	io.WriteString(w, `{"message": "Não entre em pânico!"}`)
}

func main() {
	http.HandleFunc("/", dontPanic)
	http.ListenAndServe(":9999", nil)
}
