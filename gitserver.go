package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
func getProphetResponse(w http.ResponseWriter, r *http.Request, request AppRequest) {
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			//extract body
			//var req prophetRequest
			//json.NewDecoder(r.Body).Decode(&req)
			// get github URL from body
			var projectUrl = request.Url
			// download github repository
			cloneRepo(githubUrl + projectUrl)
			//extract the absolute path
			var absolutePath = tmpServerPath + getRepoName(projectUrl)

			communicationChan := make(chan Communication)
			contextMapChan := make(chan ContextMap)

			go postProphetCommunication(communicationChan)
			go postProphetContextMap(contextMapChan)

			// post prophet for communication
			var r *http.Response = postProphet(absolutePath)
			defer r.Body.Close()
			var p ProphetResponse
			json.NewDecoder(r.Body).Decode(&p)
			//logger(w, "Sending request")
			// prophet response to json
			js, err := json.Marshal(p)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}


			//post prophet for context map

			//combine
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else {
			//request exhausted
			var errText string = "Resources exhausted, next available will be tomorrow"
			logger(w, errText)
			http.Error(w, errText, http.StatusBadRequest)
		}
	} else {
		curr = 0
		reqDate = time.Now()
		logger(w, "Sending request")
	}
}

func logger(w http.ResponseWriter, str string) {
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
	buffer := bytes.NewBuffer(newRequest(url))
	response , err := http.Post(prophetUrl,"application/json", buffer)
	if err != nil {
		panic(err)
	}
	return response
}

type ProphetRequest struct {
	Url string `json:"url"`
}


func newRequest(url string) []byte {
	var r ProphetRequest
	r.Url = url
	//r := AppRequest{Url: Url}
	requestBody, err := json.Marshal(r)
	if err != nil {

	}
	fmt.Println(string(requestBody))
	return requestBody
}

func getRepoName(githubUrl string) string{
	var s = strings.Split(githubUrl, "/")
	return s[len(s)-1]
}


// model

//type prophetRequest struct {
//	Url string
//}
type ContextMapChanStruct struct {
	PathToRepository string
	ContextMap ContextMap
}

type ContextMap struct {

}

type CommunicationChanStruct struct {

}


type ProphetResponse struct {
	Communication Communication
	ContextMap []string
}

type Communication struct {
	Edges []Edge
	Nodes []Node
}

type Edge struct {
	IdA string
	IdB string
}

type Node struct {
	Id string
}

