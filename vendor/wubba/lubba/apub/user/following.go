package user

import (
	"net/http"
	"wubba/lubba/apub"
)

const FollowingPath = "/apub/following"

type FollowingHandler struct {
	apub.BaseHandler
	GetFollowing func(string) ([]apub.User, error)
}

func (h *FollowingHandler) RoutePath() string {
	return FollowingPath
}

func (h *FollowingHandler) ServeUser(info apub.User, w http.ResponseWriter, r *http.Request) {
}
