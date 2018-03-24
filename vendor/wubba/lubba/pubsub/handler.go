package pubsub

import (
	"net/http"
	"wubba/lubba/apub"
)

type Transport struct {
	Finder func(string) (apub.User, error)
}

func (h *Transport) Broadcast(post apub.Post) (err error) {
	return
}

func (h *Transport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
