package main

import (
	"net/http"
	"os"
)

func main() {
	port := "8081"
	portConfig := os.Getenv("HTTP_PLATFORM_PORT")
	if portConfig != "" {
		port = portConfig
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":"+port, nil)
}
