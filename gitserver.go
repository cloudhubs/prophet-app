package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var MaxRequests = 5

var curr = 0
var reqDate = time.Now()
var currentTime = time.Now()
var format = "2006.01.02 15:04:05"
var p = fmt.Fprintf
var prophetUrl = "http://127.0.0.1:8081/analyze"
var tmpServerPath = "~/tmp/"
var githubUrl = "https://github.com/"

//ToDo: param
func GitServer(w http.ResponseWriter, r *http.Request) {
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			//extract body
			var req prophetRequest
			json.NewDecoder(r.Body).Decode(&req)
			// get github URL from body
			var projectUrl = req.url
			// download github repository
			cloneRepo(githubUrl + projectUrl)
			//extract the absolute path
			var absolutePath = tmpServerPath + getRepoName(projectUrl)
			// post prophet
			var r *http.Response = postProphet(absolutePath)
			defer r.Body.Close()
			var p ProphetResponse
			json.NewDecoder(r.Body).Decode(&p)
			log(w, "Sending request")
			// prophet response to json
			js, err := json.Marshal(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
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

func cloneRepo(repo string){
	cmd := exec.Command("/bin/sh", "-c", "cd " + tmpServerPath + "; git clone " + repo + ";")
	err := cmd.Run()
	if err != nil {
		// something went wrong
	}
}

func postProphet(url string) *http.Response {
	response , err := http.Post(prophetUrl,"application/json", bytes.NewBuffer(newRequest(url)) )
	if err != nil {
		panic(err)
	}
	return response
}

func newRequest(url string) []byte {
	r := prophetRequest{url: url}
	requestBody, err := json.Marshal(r)
	if err != nil {

	}
	return requestBody
}

func getRepoName(githubUrl string) string{
	var s = strings.Split(githubUrl, "/")
	return s[len(s)-1]
}


// model

type prophetRequest struct {
	url string
}

type ProphetResponse struct {
	communication Communication
	contextMap []string
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

