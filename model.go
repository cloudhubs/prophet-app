package main

type ProphetWebRequest struct {
	Url string `json:"url"`
}

type ProphetAppRequest struct {
	Path string `json:"path"`
	Cached bool `json:"cached"`
	PersistDb bool `json:"persistDb"`
	IsMonolith bool `json:"isMonolith"`
	All bool `json:"all"`
	Communication bool `json:"communication"`
}

type ProphetAppData struct {
	Global Global `json:"global"`
	Ms []Ms `json:"ms"`
}

type Global struct {
	ProjectName string `json:"projectName"`
	Communication string `json:"communication"`
	ContextMap string `json:"contextMap"`
}

type Ms struct {
	Name string `json:"name"`
	BoundedContext string `json:"boundedContext"`
}