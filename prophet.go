package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var prophetAPIString = "http://127.0.0.1:8081/"


func queryProphet(w http.ResponseWriter, r *http.Request) {

	// if check.resources OK
	cont, msg := checkCapacity()
	if !cont {
		sendForbidden(w, msg)
	}

	// unmarshall
	webAppReq, err := unmarshallRequest(r)
	if err != nil {
		sendServerError(w, err.Error())
	}

	//noinspection GoNilness
	for _, s := range webAppReq.Repositories {
		// if check.orgRepo OK
		cont, msg = checkGit(s.Organization, s.Repository)
		if !cont {
			sendNotFound(w, msg)
		}
	}

	// make a query
	requestBody, err := json.Marshal(webAppReq)
	buffer := bytes.NewBuffer(requestBody)
	response, err := fetchProphet(buffer)
	if err != nil {
		sendServerError(w, err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	_, err = w.Write(body)
	if err != nil {
		sendServerError(w, err.Error())
	}
}

func unmarshallRequest(r http.Request) (ProphetWebRequest, error) {
	err := r.Body.Close()
	if err != nil {
		return ProphetWebRequest{}, err
	}
	var p ProphetWebRequest
	err = json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		return ProphetWebRequest{}, err
	}
	return p, nil
}


func sendForbidden(w http.ResponseWriter, reason string){
	logger(w, reason)
	http.Error(w, reason, http.StatusForbidden)
}

func sendNotFound(w http.ResponseWriter, reason string){
	logger(w, reason)
	http.Error(w, reason, http.StatusNotFound)
}

func sendServerError(w http.ResponseWriter, reason string){
	logger(w, reason)
	http.Error(w, reason, http.StatusInternalServerError)
}

func postProphetAPI(buffer *bytes.Buffer) *http.Response {
	response , err := http.Post(prophetAPIString,"application/json", buffer)
	if err != nil {
		panic(err)
	}
	return response
}