package main

import (
	"testing"
)




//func TestMainEndpoint(t *testing.T){
//	//Url := "/cloudhubs/tms"
//
//	//bytes.NewBuffer(newRequest(Url) )
//	resp, err := http.Post("http://127.0.0.1:8080","json", nil )
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	var p ProphetResponse
//	json.NewDecoder(resp.Body).Decode(&p)
//	if p.communication.edges == nil {
//		t.Errorf("Communication edges, got: %d, want more then: %d.", len(p.communication.edges), 0)
//	}
//
//}


func TestCloneRepo(t *testing.T){
	cloneRepo("https://github.com/cloudhubs/prophet")
}

