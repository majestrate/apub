package outbox

import (
	"net/http"
	"wubba/lubba/apub"
)

const RoutePath = "/apub/outbox"

type Handler struct {
	Finder func(string) (apub.UserInfo, error)
}

func (h *Handler) RoutePath() string {
	return RoutePath
}

func (h *Handler) ServeUser(info apub.UserInfo, w http.ResponseWriter, r *http.Request) {
}
