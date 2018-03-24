package apub

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"io"
	"strings"
	"time"
	"wubba/lubba/apub"
	"wubba/lubba/apub/atom"
	"wubba/lubba/apub/model"
	"wubba/lubba/apub/util"
)

type User = apub.User

// UserName is the @user@host
type UserName string

func (u UserName) String() string {
	return string(u)
}

// user part
func (u UserName) User() string {
	parts := strings.Split(u.String(), "@")
	if len(parts) == 3 {
		return parts[1]
	}
	return ""
}

// NormalizeUser converts a string into a UserName given our local server name
func NormalizeUser(str, ourserver string) UserName {
	if str[0] != '@' {
		str = "@" + str
	}
	u := UserName(str)
	if u.Server() == "" {
		u = UserName(str + "@" + ourserver)
	}
	return u
}

// server part
func (u UserName) Server() string {
	parts := strings.Split(u.String(), "@")
	if len(parts) == 3 && len(parts[2]) > 0 {
		return parts[2]
	}
	return ""
}

// implements apub.User
type UserInfo struct {
	ServerName   string // server name
	UserName     string // username without '@'
	PreferedName string // prefered name

	Summary string // profile summary
	Avatar  string // avatar url
	Header  string // header url

	SigningKey *rsa.PrivateKey // rsa private signing key

	Profile string // profile url
	Inbox   string // inbox url
	Outbox  string // outbox url
	Feed    string // atom feed url

	GetPosts       func(offset int64, limit int) ([]*Post, error)
	GetLastUpdated func() time.Time
}

func (info *UserInfo) Posts(offset int64, limit int) (posts []apub.Post, err error) {
	if info.GetPosts != nil {
		var localposts []*Post
		localposts, err = info.GetPosts(offset, limit)
		if len(localposts) > 0 {
			posts = make([]apub.Post, len(localposts))
			for idx := range localposts {
				posts[idx] = localposts[idx]
			}
		}
	}
	return
}

func (info *UserInfo) RegenerateSigningKey() (err error) {
	info.SigningKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err == nil {
		info.SigningKey.Precompute()
	}
	return
}

func (info *UserInfo) User() string {
	return "@" + info.UserName + "@" + info.ServerName
}

func (info *UserInfo) WebFingerAlias() string {
	return info.ProfileURL()
}

func (info *UserInfo) WebFingerSubject() string {
	return "acct:" + info.User()[1:]
}

func (info *UserInfo) ProfileURL() string {
	return info.Profile
}

func (info *UserInfo) AtomFeedURL() string {
	return info.Feed
}

func (info *UserInfo) InboxURL() string {
	return info.Inbox
}

func (info *UserInfo) OutboxURL() string {
	return info.Outbox
}

func (info *UserInfo) LastUpdated() time.Time {
	if info.GetLastUpdated == nil {
		return time.Now()
	}
	return info.GetLastUpdated()
}

func (info *UserInfo) ToAtomFeed(title string, nextURL string) (f apub.UserFeed, err error) {
	feedURL := info.AtomFeedURL()
	profile := info.ProfileURL()
	updated := info.LastUpdated()
	feedLinks := []apub.Link{
		apub.Link{
			Rel:  apub.SelfRel,
			Href: feedURL,
			Type: apub.AtomMime,
		},
	}
	if nextURL != "" {
		feedLinks = append(feedLinks, apub.Link{
			Rel:  apub.NextRel,
			Href: nextURL,
			Type: apub.AtomMime,
		})
	}
	f = &atom.Feed{
		Published: updated,
		Updated:   updated,
		Logo:      info.Avatar,
		Links:     feedLinks,
		Author: model.Author{
			ID:           profile,
			DisplayName:  info.PreferedName,
			PreferedName: info.PreferedName,
			Name:         info.UserName,
			AP:           true,
			Note:         info.Summary,
			Summary:      info.Summary,
			ObjectType:   model.TypePerson,
			URI:          profile,
			Links: []apub.Link{
				apub.Link{
					Rel:  "avatar",
					Href: info.Avatar,
				},
				apub.Link{
					Rel:  "header",
					Href: info.Header,
				},
			},
		},
		ID:    feedURL,
		Title: title,
	}
	return
}

// sign data using rsa / sha256
func (info *UserInfo) Sign(r io.Reader) (sig, hash []byte, err error) {
	var buff [1024]byte
	h := sha256.New()
	io.CopyBuffer(h, r, buff[:])
	shahash := h.Sum(nil)
	hash = shahash[:]
	sig, err = info.SigningKey.Sign(rand.Reader, hash, crypto.SHA256)
	return
}

func (info *UserInfo) WebFingerLinks() (links []apub.Link) {
	links = make([]apub.Link, 4)
	links[0] = apub.Link{
		Rel:  apub.AtomRel,
		Type: apub.AtomMime,
		Href: info.AtomFeedURL(),
	}
	links[1] = apub.Link{
		Rel:  apub.WebFingerRel,
		Type: apub.HTMLMime,
		Href: info.ProfileURL(),
	}
	links[2] = apub.Link{
		Rel:  apub.MagicKeyRel,
		Type: apub.MagicKeyMime,
		Href: "data:" + apub.MagicKeyMime + "," + util.EncodeRSAPublicKey(info.SigningKey.PublicKey),
	}
	links[3] = apub.Link{
		Rel:  apub.SelfRel,
		Href: info.ProfileURL(),
		Type: apub.ActivityMime,
	}
	return
}
