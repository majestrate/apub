package apub

import (
	"net/http"
	"wubba/lubba/apub/hostmeta"
	"wubba/lubba/apub/webfinger"
)

// APubHandler
type APubHandler struct {
	finger   webfinger.WebFinger
	hostmeta hostmeta.Handler
}

// Setup sets up routes
func (a *APubHandler) Setup(finder func(string) (UserInfo, error), setupRoute func(string, http.Handler)) {
	a.finger.Finder = finder
	setupRoute("/.well-known/host-meta", &a.hostmeta)
	setupRoute("/.well-known/webfinger", &a.finger)
	return
}
