package main

import (
	"fmt"
	"net/http"
)
var format = "2006.01.02 15:04:05"
var p = fmt.Fprintf


func logger(w http.ResponseWriter, str string) {
	err, _ := p(w, str + " %d, %s, %s", curr, reqDate.Format(format), currentTime.Format(format))
	if err == 0 {
		_, _ = p(w, "Error.")
	}
}
