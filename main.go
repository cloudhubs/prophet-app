package main

import (
	"errors"
	"github.com/rs/cors"
	_ "github.com/rs/cors"
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
	js := getProphetResponse(w,r,p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func analyzeOptions(w http.ResponseWriter, r *http.Request) {
	log.Println("Options...")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}


func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", analyzeGit)
	//log.Println("Starting server on :8080...")

	//router := mux.NewRouter()
	//router.HandleFunc("/", analyzeGit)
	//log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

	mux := http.NewServeMux()
	mux.HandleFunc("/", analyzeGit)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods: []string{"HEAD","GET","POST"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Content-Type", "Accept"},
		//ContentType: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// Insert the middleware
	handler = c.Handler(handler)

	http.ListenAndServe(":8080", handler)

	//router := NewRouter()
	//log.Fatal(http.ListenAndServe(":3000", s.CORS(s.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), s.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), s.AllowedOrigins([]string{"*"}))(router)))


	//r := NewRouter()
	//r.HandleFunc("/", analyzeGit)//.Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch)
	//r.HandleFunc("/", analyzeOptions).Methods(http.MethodOptions)

	//r.Use(CORSMethodMiddleware(r))
	//http.ListenAndServe(":8080", r)

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