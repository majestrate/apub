package webfinger

import (
	"net/http"
	"wubba/lubba/apub"
)

type WebFinger struct {
	Finder apub.UserFinder
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Response struct {
	Subject string        `json:"subject"`
	Links   []interface{} `json:"links"`
}

func (wf *WebFinger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if wf.Finder == nil {
		// no finder
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resource := r.URL.Query().Get("resource")
	u, err := wf.Finder(resource)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if u == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
