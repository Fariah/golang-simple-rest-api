package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%s running at 127.0.0.1 on port 8080", "server")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}
}
