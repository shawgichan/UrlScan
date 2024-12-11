package main

import (
	"UrlScan/internal"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := internal.NewHandler()
	mux.Handle("/scan", handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
