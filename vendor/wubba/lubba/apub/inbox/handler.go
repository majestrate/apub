package inbox

import (
	"net/http"
	"wubba/lubba/apub"
)

const RoutePath = "/apub/inbox"

type Handler struct {
	apub.BaseHandler
}

func (h *Handler) RoutePath() string {
	return RoutePath
}

func (h *Handler) ServeUser(info apub.UserInfo, w http.ResponseWriter, r *http.Request) {
}
