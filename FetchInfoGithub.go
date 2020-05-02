package main

import (
	"bytes"
	"net/http"
)


var mainGithubAPI = "http://github.com/api/"

/**
Fetch Organization to validate existence
*/
func fetchOrg(buffer *bytes.Buffer) (*http.Response, error) {
	response , err := http.Get(prophetAPI)
	return response, err
}

/**
Fetch Repository from Organization to validate existence
 */
func fetchOrgRepo(org string, repo string) (*http.Response, error) {
	url := mainGithubAPI + org + "/" + repo
	response , err := http.Get(url)
	return response, err
}
