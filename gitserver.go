package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
	"os"
)

var MaxRequests = 500
//var curr = 0
//var reqDate = time.Now()
//var currentTime = time.Now()
var format = "2006.01.02 15:04:05"
var p = fmt.Fprintf
var prophetAPIString = "http://127.0.0.1:8081/multirepoapp"
var tmpServerPath = "/Users/svacina/tmp/"
var githubUrl = "https://github.com/"


func getRepoName(githubUrl string) string{
	var s = strings.Split(githubUrl, "/")
	return s[len(s)-1]
}

//ToDo ResponseWriter
func createProphetRequest(arr []string, monoliths []bool) []byte {
	var reqs []ProphetAppRequest
	for i:= 0 ; i < len(arr); i++ {
		var r ProphetAppRequest
		r.Path = arr[i]
		r.IsMonolith = monoliths[i]
		reqs = append(reqs, r)
	}
	var multi ProphetAppMultiRepoRequest
	multi.Repositories = reqs
	multi.SystemName = "default"
	requestBody, err := json.Marshal(multi)
	if err != nil {
		//ToDo send error
	}
	fmt.Println(string(requestBody))
	return requestBody
}

//ToDo ResponseWriter
func cloneRepo(repo string){
	cmd := exec.Command("/bin/sh", "-c", "cd " + tmpServerPath + "; git clone " + repo + ";")
	err := cmd.Run()
	if err != nil {
		// ToDo send error
	}
}


//ToDo ResponseWriter
func deleteRepo(repo string){
	if repo != "/" {
		os.RemoveAll(tmpServerPath + repo)
	}


}

func postProphetAPI(buffer *bytes.Buffer) *http.Response {
	response , err := http.Post(prophetAPIString,"application/json", buffer)
	if err != nil {
		panic(err)
	}
	return response
}

func getProphetAppData(r *http.Response) ProphetAppData{
	defer r.Body.Close()
	var p ProphetAppData
	json.NewDecoder(r.Body).Decode(&p)
	return p
}

func marshalProphetAppData(p ProphetAppData) ([]byte, error) {
	js, err := json.Marshal(p)
	if err != nil {
		//ToDo
	}
	return js, err
}

func callProphet(request ProphetWebRequest) ([]byte, error) {
	//init array
	var arr []string
	var monoliths []bool
	for i := 0; i < len(request.Repositories); i++ {
		var projectUrl = request.Repositories[i].Organization + "/" + request.Repositories[i].Repository;
		cloneRepo(githubUrl + projectUrl)
		repoName := getRepoName(projectUrl)
		var absolutePath = tmpServerPath + repoName
		arr = append(arr, repoName)
		arr[i] = absolutePath //put absolute path to an array
		monoliths = append(monoliths, request.Repositories[i].IsMonolith)
	}
	buffer := bytes.NewBuffer(createProphetRequest(arr, monoliths))
	//var projectUrl = request.Url
	//cloneRepo(githubUrl + projectUrl)
	//repoName := getRepoName(projectUrl)
	//var absolutePath = tmpServerPath + repoName
	//buffer := bytes.NewBuffer(createProphetRequest(absolutePath))
	response := postProphetAPI(buffer)
	prophetAppData := getProphetAppData(response)
	//we have the data and we can delete
	//for i := 0; i < len(arr); i++ {
	//	repoName := getRepoName(arr[i])
	//	deleteRepo(repoName)
	//}
	return marshalProphetAppData(prophetAppData)
}

func logger(w http.ResponseWriter, str string) {
	err, _ := p(w, str + " %d, %s, %s", curr, reqDate.Format(format), currentTime.Format(format))
	if err == 0 {
		p(w, "Error.")
	}
}


func getProphetResponse(w http.ResponseWriter, r *http.Request, request ProphetWebRequest) ([]byte, error) {
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			return callProphet(request)
		} else {
			//request exhausted
			var errText = "Resources exhausted, next available will be tomorrow"
			logger(w, errText)
			http.Error(w, errText, http.StatusBadRequest)
		}
	}
	curr = 0
	reqDate = time.Now()
	return callProphet(request)

}
