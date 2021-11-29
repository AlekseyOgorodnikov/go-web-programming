package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
type WorldHandler struct{}
type AllHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func (h *AllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Aleksey in new world!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	all := AllHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	http.Handle("/", &all)

	server.ListenAndServe()
}
