package user

import (
	"net/http"
	"wubba/lubba/apub"
)

const FollowingPath = "/apub/following"

type FollowingHandler struct {
	apub.BaseHandler
	GetFollowing func(string) ([]apub.UserInfo, error)
}

func (h *FollowingHandler) RoutePath() string {
	return FollowingPath
}

func (h *FollowingHandler) ServeUser(info apub.UserInfo, w http.ResponseWriter, r *http.Request) {
}
