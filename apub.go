package apub

import (
	"net/http"
	"wubba/lubba/apub"
	"wubba/lubba/apub/hostmeta"
	"wubba/lubba/apub/webfinger"
)

// Router implements http.Handler
type Router struct {
	mux      *http.ServeMux
	finger   webfinger.WebFinger
	hostmeta *hostmeta.Handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}

// create a new router given a URLFinder that handles webfinger
func NewRouter(finder URLFinder, hostname string) (r *Router) {
	r = new(Router)
	r.finger.Finder = apub.UserFinder(finder)
	r.hostmeta = hostmeta.NewHandler(hostname)
	r.mux = http.NewServeMux()
	r.mux.Handle("/.well-known/host-meta", r.hostmeta)
	r.mux.Handle("/.well-known/webfinger", &r.finger)
	return
}
