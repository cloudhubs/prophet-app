package main

import (
	"net/http"
)

func main() {
    http.HandleFunc("/", GitServer)
    http.ListenAndServe(":8080", nil)
}