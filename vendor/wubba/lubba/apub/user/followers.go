package user

import (
	"net/http"
	"wubba/lubba/apub"
)

const FollowersPath = "/apub/followers"

type FollowersHandler struct {
	apub.BaseHandler
	GetFollowers func(string) ([]apub.User, error)
}

func (h *FollowersHandler) RoutePath() string {
	return FollowersPath
}

func (h *FollowersHandler) ServeUser(info apub.User, w http.ResponseWriter, r *http.Request) {
}
