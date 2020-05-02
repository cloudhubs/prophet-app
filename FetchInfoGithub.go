package main

import (
	"net/http"
)


var orgRepoAPI = "http://api.github.com/repos/"
var orgAPI = "http://api.github.com/orgs/"

/**
Fetch Organization to validate existence
*/
func fetchOrg(org string) (*http.Response, error) {
	response , err := http.Get(orgAPI + "/" + org)
	return response, err
}

/**
Fetch Repository from Organization to validate existence
 */
func fetchOrgRepo(org string, repo string) (*http.Response, error) {
	url := orgRepoAPI + org + "/" + repo
	response , err := http.Get(url)
	return response, err
}
