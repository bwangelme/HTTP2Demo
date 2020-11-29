package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello, h2c")
	})
	h2s := http2.Server{}
	h1s := &http.Server{
		Addr:    ":8972",
		Handler: h2c.NewHandler(handler, &h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}
