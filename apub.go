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

	Database Database // Database must not be nil
}

// SetupRoutes sets up routes
func (a *APubHandler) SetupRoutes(setupRoute func(string, http.Handler), setupSubRouter func(string, http.Handler)) {
	// set up finder
	localfinder := func(str string) (User, error) {
		if a.Database == nil {
			return nil, nil
		}
		user := NormalizeUser(str, a.Database.LocalHost())
		return a.Database.LocalUser(user.User())
	}
	a.finger.Finder = localfinder
	a.inbox.Finder = localfinder
	a.outbox.Finder = localfinder
	a.followers.Finder = localfinder
	a.following.Finder = localfinder
	a.feeds.Finder = localfinder

	a.followers.GetFollowers = func(str string) (infos []User, err error) {
		if a.Database != nil {
			infos, err = a.Database.ListFollowers(str)
		}
		return infos, err
	}
	a.following.GetFollowing = func(str string) (infos []User, err error) {
		if a.Database != nil {
			infos, err = a.Database.ListFollowing(str)
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
