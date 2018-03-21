package apub

import (
	"net/http"
	"wubba/lubba/apub"
	"wubba/lubba/apub/webfinger"
)

// Router implements http.Handler
type Router struct {
	mux    *http.ServeMux
	finger webfinger.WebFinger
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

// create a new router given a URLFinder that handles webfinger
func NewRouter(finder URLFinder) (r Router) {
	r.finger.Finder = apub.UserFinder(finder)
	r.mux = http.NewServeMux()
	r.mux.Handle("/.well-known/webfinger", &r.finger)
	return
}
