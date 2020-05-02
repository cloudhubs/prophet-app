package main

func checkGit(org string, repo string) (bool, string){
	_, err := fetchOrg(org)
	if err != nil {
		return false, "Organization not found"
	}

	_, err = fetchOrgRepo(org, repo)
	if err != nil {
		return false, "Repository not found"
	}

	return true, ""
}
