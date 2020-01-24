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
var prophetUrl = "http://127.0.0.1:8081/"
var communicationInterface = "communication"
var contextMapInterface = "contextmap"
var tmpServerPath = "/Users/svacina/tmp/"
var githubUrl = "https://github.com/"

type CurrentRequest struct {
	Value int `json:"value"`
}

func getCurrentRequestValue(w http.ResponseWriter, r *http.Request) {
	cr := CurrentRequest{
		Value: curr,
	}
	js, err := json.Marshal(cr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}



//ToDo: param
func getProphetResponse(w http.ResponseWriter, r *http.Request, request AppRequest) {
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			var projectUrl = request.Url
			cloneRepo(githubUrl + projectUrl)

			// chan init
			communicationChan := make(chan CommunicationChan)
			contextMapChan := make(chan ContextMapChan)

			var absolutePath = tmpServerPath + getRepoName(projectUrl)

			// routines
			go postProphetCommunication(communicationChan)
			go postProphetContextMap(contextMapChan)

			// send objects to channel
			var ccm = ContextMapChan{
				PathToRepository: absolutePath,
				ContextMap:       ContextMap{},
			}
			contextMapChan <- ccm

			var cmc = CommunicationChan{
				PathToRepository: absolutePath,
				Communication:    Communication{},
			}

			communicationChan <- cmc

			ccm, cmc = <-contextMapChan, <-communicationChan

			//var r *http.Response = postProphet(absolutePath)
			//defer r.Body.Close()
			//var p ProphetResponse
			//json.NewDecoder(r.Body).Decode(&p)
			pr := ProphetResponse{
				Communication: cmc.Communication,
				ContextMap:    ccm.ContextMap,
			}
			js, err := json.Marshal(pr)
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

func postProphetContextMap(c chan ContextMapChan) {
	obj := <- c
	var r = postProphet(obj.PathToRepository, contextMapInterface)
	defer r.Body.Close()
	var p ContextMap
	json.NewDecoder(r.Body).Decode(&p)
	obj.ContextMap = p
	c <- obj
}

func postProphetCommunication(c chan CommunicationChan) {
	obj := <- c
	var r *http.Response = postProphet(obj.PathToRepository, communicationInterface)
	defer r.Body.Close()
	var p Communication
	json.NewDecoder(r.Body).Decode(&p)
	obj.Communication = p
	c <- obj
}

func postProphet(url string, pathInterface string) *http.Response {
	buffer := bytes.NewBuffer(newRequest(url))
	response , err := http.Post(prophetUrl + pathInterface,"application/json", buffer)
	if err != nil {
		panic(err)
	}
	return response
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

// channels
type ContextMapChan struct {
	PathToRepository string `json:"pathToRepository"`
	ContextMap ContextMap `json:"contextMap"`
}

type CommunicationChan struct {
	PathToRepository string
	Communication Communication
}

// model

type ContextMap struct {
	MarkdownStrings []string `json:"markdownStrings"`
}

type Communication struct {
	Edges []Edge `json:"edges"`
	Nodes []Node `json:"nodes"`
}

type Edge struct {
	From string `json:"from"`
	To string `json:"to"`
}

type Node struct {
	Id string `json:"id"`
	Label string `json:"label"`
}

type ProphetResponse struct {
	Communication Communication `json:"communication"`
	ContextMap ContextMap `json:"contextMap"`
}





