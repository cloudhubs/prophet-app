package main

import (
	"bytes"
	"net/http"
)
//ToDo: change to docker variable
var prophetAPI = "http://127.0.0.1:8081/"

/**
Fetch Data from Prophet Utils App Endpoint
 */
func fetchProphet(buffer *bytes.Buffer) (*http.Response, error) {
	response , err := http.Post(prophetAPI,"application/json", buffer)
	return response, err
}
