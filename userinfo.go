package apub

import "wubba/lubba/apub"

// implements apub.UserInfo
type UserInfo struct {
	ServerName string
	UserName   string
	PublicKey  string
	ProfileURL string
	AtomURL    string
	InboxURL   string
	OutboxURL  string
}

func (info *UserInfo) Alias() string {
	return info.ProfileURL
}

func (info *UserInfo) Subject() string {
	return "acct:" + info.UserName + "@" + info.ServerName
}

func (info *UserInfo) Links() (links []apub.Link) {
	links = make([]apub.Link, 5)
	links[0] = apub.Link{
		Rel:  apub.AtomRel,
		Type: apub.AtomMime,
		Href: info.AtomURL,
	}
	links[1] = apub.Link{
		Rel:  apub.WebFingerRel,
		Type: apub.HTMLMime,
		Href: info.ProfileURL,
	}
	return
}
