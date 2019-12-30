//func main() {
//    http.HandleFunc("/", GitServer)
//    http.ListenAndServe(":8080", nil)
//}

package main

import (
	"errors"
	"log"
	"net/http"
)

type AppRequest struct {
	Url string
}

func analyzeGit(w http.ResponseWriter, r *http.Request) {
	var p AppRequest

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
	//analyze
	getProphetResponse(w,r,p)
	//fmt.Fprintf(w, "AppRequest: %+v", p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", analyzeGit)
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}