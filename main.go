package main

import (
	"errors"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
	"log"
	"net/http"
)



func analyzeGit(w http.ResponseWriter, r *http.Request) {
	log.Println("Analyzing...")

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	js := getProphetResponse(w,r,p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", analyzeGit)
	//log.Println("Starting server on :8080...")
	r := mux.NewRouter()
	r.HandleFunc("/", analyzeGit).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", r)

	//c := cors.New(cors.Options{
	//	AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
	//	AllowCredentials: true,
	//	// Enable Debugging for testing, consider disabling in production
	//	Debug: true,
	//})

	//// Insert the middleware
	//handler = c.Handler(handler)
	//mux.Handler(c)




	//handler := cors.Default().Handler(mux)

	//err := http.ListenAndServe(":8080", handler)
	//log.Fatal(err)
}