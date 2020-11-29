package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute

func main() {
	var srv http.Server
	srv.Addr = ":8972"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, http/2")
	})

	err := http2.ConfigureServer(&srv, &http2.Server{})
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Fatal(srv.ListenAndServeTLS(
			"/home/xuyundong/Certs/httpscert.pem", "/home/xuyundong/Certs/httpskey.pem",
		))
	}()
	select {}
}
