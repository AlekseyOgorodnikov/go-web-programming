package main

import (
	"fmt"
	"net/http"
)

// handler path function input-output
func handler(writer http.ResponseWriter, requets *http.Request) {
	fmt.Fprintf(writer, "Hello world, Aleksey, %s!", requets.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
