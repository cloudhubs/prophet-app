package main

import (
	"time"
)

var curr = 0
var reqDate = time.Now()
var currentTime = time.Now()

func checkCapacity() (bool, string){
	curr = curr + 1
	currentTime = time.Now()
	diff := currentTime.Sub(reqDate)
	if diff.Hours() < 24 {
		if curr < MaxRequests {
			return true, ""
		} else {
			return false, "Capacity fulfilled. Check tomorrow."
		}
	}
	curr = 0
	reqDate = time.Now()
	return true, ""
}