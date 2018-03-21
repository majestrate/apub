package webfinger

import (
	"net/http"
	"wubba/lubba/apub"
)

type WebFinger struct {
	Finder apub.UserFinder
}

func (wf *WebFinger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
