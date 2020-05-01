package main

import (
	"errors"
	"github.com/rs/cors"
	"log"
	"net/http"
	)


func analyzeGit(w http.ResponseWriter, r *http.Request) {
	log.Println("Analyzing...")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	js, err := getProphetResponse(w,r,p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/**
Main server function
 */
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", analyzeGit)
	handler := cors.Default().Handler(mux)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods: []string{"HEAD","GET","POST"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Accept"},
		Debug: true,
	})
	handler = c.Handler(handler)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Println("Error starting server at 8080" + err.Error())
	}
}