package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
)

/**
Main server function
 */
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", queryProphet)
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