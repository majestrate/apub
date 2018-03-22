package apub

import (
	"net/http"
	"wubba/lubba/apub"
	"wubba/lubba/apub/atom"
	"wubba/lubba/apub/federation"
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
	feeds     atom.Handler
	federator federation.Federator

	Database Database // Database implements UserFinder
}

// SetupRoutes sets up routes
func (a *APubHandler) SetupRoutes(setupRoute func(string, http.Handler), setupSubRouter func(string, http.Handler)) {
	// set up finder
	localfinder := func(str string) (apub.UserInfo, error) {
		if a.Database == nil {
			return nil, nil
		}
		return a.Database.LocalUser(str)
	}
	a.finger.Finder = localfinder
	a.inbox.Finder = localfinder
	a.outbox.Finder = localfinder
	a.followers.Finder = localfinder
	a.following.Finder = localfinder
	a.feeds.Finder = localfinder

	a.followers.GetFollowers = func(str string) (infos []apub.UserInfo, err error) {
		if a.Database != nil {
			var users []*UserInfo
			users, err = a.Database.ListFollowers(str)
			if err == nil {
				if len(users) > 0 {
					infos = make([]apub.UserInfo, len(users))
					for idx := range users {
						infos[idx] = users[idx]
					}
				}
			}
		}
		return infos, err
	}
	a.following.GetFollowing = func(str string) (infos []apub.UserInfo, err error) {
		var users []*UserInfo
		if a.Database != nil {
			users, err = a.Database.ListFollowing(str)
			if err == nil {
				if len(users) > 0 {
					infos = make([]apub.UserInfo, len(users))
					for idx := range users {
						infos[idx] = users[idx]
					}
				}
			}
		}
		return infos, err
	}

	// set up routes
	setupRoute("/.well-known/host-meta", &a.hostmeta)
	setupRoute("/.well-known/webfinger", &a.finger)

	handlers := []apub.UserHandler{&a.inbox, &a.outbox, &a.followers, &a.following, &a.feeds}

	// set up handlers
	for idx := range handlers {
		setupSubRouter(handlers[idx].RoutePath(), &apub.BaseHandler{
			Finder:      localfinder,
			UserHandler: handlers[idx],
		})
	}
}
