package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	handler := HelloHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
