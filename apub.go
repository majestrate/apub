package apub

import (
	"net/http"
	"wubba/lubba/apub/webfinger"
)

// Router implements http.Handler
type Router struct {
	mux    *http.ServeMux
	finger webfinger.WebFinger
}

func (r *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.mux.ServeHTTP(w, r)
}

// create a new router given a UserFinder
func NewRouter(finder UserFinder) (r Router) {
	r.finger.Finder = finder
	r.mux = http.NewServeMux()
	r.mux.Handle("/.well-known/webfinger", r.finger)
	return
}
