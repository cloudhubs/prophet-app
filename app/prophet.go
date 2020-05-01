package app

import (
	"net/http"
	"time"
)

func queryProphet() {
	// if check.resources OK

	// if check.orgRepo OK

	// make a query
}

var curr = 0
var reqDate = time.Now()
var currentTime = time.Now()

func checkResources() bool{
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			return true
		} else {
			return false
		}
	}
	curr = 0
	reqDate = time.Now()
	return true
}

func sendResourcesExhausted(w http.ResponseWriter){
	var errText = "Resources exhausted, next available will be tomorrow"
	logger(w, errText)
	http.Error(w, errText, http.StatusBadRequest)
}

func sendRepoNotOK(w http.ResponseWriter, reason string){
	var errText = "Error in repository: " + reason
	logger(w, errText)
	http.Error(w, errText, http.StatusBadRequest)
}
