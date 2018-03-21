package apub

import (
	"net/http"
	"wubba/lubba/apub"
	"wubba/lubba/apub/hostmeta"
	"wubba/lubba/apub/inbox"
	"wubba/lubba/apub/outbox"
	"wubba/lubba/apub/user"
	"wubba/lubba/apub/webfinger"
)

// APubHandler
type APubHandler struct {
	finger    webfinger.WebFinger
	hostmeta  hostmeta.Handler
	inbox     inbox.Handler
	outbox    outbox.Handler
	followers user.FollowersHandler
	following user.FollowingHandler

	Finder InfoFinder // Finder is responsible for fetching user information
}

// Setup sets up routes and gives it a InfoFinder
func (a *APubHandler) Setup(setupRoute func(string, http.Handler), setupSubRouter func(string, http.Handler)) {
	// set up finder

	localfinder := func(str string) (apub.UserInfo, error) {
		return a.Finder.LocalUser(str)
	}
	a.finger.Finder = localfinder
	a.inbox.Finder = localfinder
	a.outbox.Finder = localfinder
	a.followers.Finder = localfinder
	a.following.Finder = localfinder

	a.followers.GetFollowers = func(str string) ([]apub.UserInfo, error) {
		var infos []apub.UserInfo
		users, err := a.Finder.ListFollowers(str)
		if err == nil {
			if len(users) > 0 {
				infos = make([]apub.UserInfo, len(users))
				for idx := range users {
					infos[idx] = users[idx]
				}
			}
		}
		return infos, err
	}
	a.following.GetFollowing = func(str string) ([]apub.UserInfo, error) {
		var infos []apub.UserInfo
		users, err := a.Finder.ListFollowing(str)
		if err == nil {
			if len(users) > 0 {
				infos = make([]apub.UserInfo, len(users))
				for idx := range users {
					infos[idx] = users[idx]
				}
			}
		}
		return infos, err
	}

	// set up routes
	setupRoute("/.well-known/host-meta", &a.hostmeta)
	setupRoute("/.well-known/webfinger", &a.finger)

	handlers := []apub.UserHandler{&a.inbox, &a.outbox, &a.followers, &a.following}

	for idx := range handlers {
		setupSubRouter(handlers[idx].RoutePath(), &apub.BaseHandler{
			Finder:      localfinder,
			UserHandler: handlers[idx],
		})
	}
}
