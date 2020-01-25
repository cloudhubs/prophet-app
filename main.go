package main

import (
	"errors"
	"log"
	"net/http"
)



func analyzeGit(w http.ResponseWriter, r *http.Request) {
	var p ProphetWebRequest
	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	js := getProphetResponse(w,r,p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", analyzeGit)
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}