package main

import (
	"encoding/json"
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
		coder := &Person{
			First: "Casey",
			Last:  "Manus",
		}
		data, _ := json.Marshal(coder)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
	http.ListenAndServe(":"+port, nil)
}

type Person struct {
	First string
	Last  string
}
