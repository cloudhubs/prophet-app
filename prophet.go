package main

import (
	"encoding/json"
	"net/http"
)

func queryProphet(r http.Request, w http.ResponseWriter) {
	// unmarshall
	prophetAppData := unmarshallData(r)
	// if check.resources OK
	cont, msg := checkCapacity()
	if !cont {
		sendForbidden(w, msg)
	}
	// if check.orgRepo OK
	cont, msg = checkGit(org, repo)
	if !cont {
		sendNotFound(w, msg)
	}
	// make a query


}

func unmarshallData(r http.Request) (ProphetAppData, error) {
	err := r.Body.Close()
	if err != nil {
		return ProphetAppData{}, err
	}
	var p ProphetAppData
	json.NewDecoder(r.Body).Decode(&p)
	return p
}


func sendForbidden(w http.ResponseWriter, reason string){
	logger(w, reason)
	http.Error(w, reason, http.StatusForbidden)
}

func sendNotFound(w http.ResponseWriter, reason string){
	logger(w, reason)
	http.Error(w, reason, http.StatusNotFound)
}
