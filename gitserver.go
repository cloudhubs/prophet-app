package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var MaxRequests = 5

var curr = 0
var reqDate = time.Now()
var currentTime = time.Now()
var format = "2006.01.02 15:04:05"
var p = fmt.Fprintf
var prophetUrl = "/"

//ToDo: param
func GitServer(w http.ResponseWriter, r *http.Request) {
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			// send request
			log(w, "Sending request")
		} else {
			//request exhausted
			log(w, "Resources exhausted, next available will be tomorrow")
		}
	} else {
		curr = 0
		reqDate = time.Now()
		log(w, "Sending request")
	}
}

func log(w http.ResponseWriter, str string) {
	err, _ := p(w, str + " %d, %s, %s", curr, reqDate.Format(format), currentTime.Format(format))
	if err == 0 {
		p(w, "Error.")
	}
}

type prophetRequest struct {
	url string
}

func postProphet(url string){
	_, err := http.Post(prophetUrl,"application/json", bytes.NewBuffer(newRequest(url)) )
	if err != nil {
		panic(err)
	}

}

func newRequest(url string) []byte {
	r := prophetRequest{url: url}
	requestBody, err := json.Marshal(r)
	if err != nil {

	}
	return requestBody
}


// model

type ProphetResponse struct {
	communication Communication
	contextMap struct{}
}

type Communication struct {
	edges []Edge
	nodes []Node
}

type Edge struct {
	idA string
	idB string
}

type Node struct {
	id string
}

