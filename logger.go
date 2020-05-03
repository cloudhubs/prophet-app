package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
var format = "2006.01.02 15:04:05"
var p = fmt.Fprintf


func logger(w http.ResponseWriter, str string) {
	errant, _ := p(w, str + " %d, %s, %s", curr, reqDate.Format(format), currentTime.Format(format))
	if errant == 0 {
		_, _ = p(w, "Error.")
	}

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println(str + " %d, %s, %s", curr, reqDate.Format(format), currentTime.Format(format))

}
