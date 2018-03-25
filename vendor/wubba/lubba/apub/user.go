package apub

import (
	"io"
	"time"
	"wubba/lubba/apub/json"
)

type UserFeed interface {
	AppendPost(p Post)
	Posts() []Post
}

// serializable user info
type User interface {
	WebFingerSubject() string
	WebFingerAlias() string
	WebFingerLinks() []Link

	ToAtomFeed(title, nextFeedURL string) (UserFeed, error)

	// User gets the @user@hostname
	User() string

	AtomFeedURL() string
	ProfileURL() string
	InboxURL() string
	OutboxURL() string

	// time last updated
	LastUpdated() time.Time

	// Sign does rsa signing on body of data
	// return signature, hash, error
	Sign(io.Reader) ([]byte, []byte, error)

	Posts(offset int64, limit int) ([]Post, error)

	json.Marshaler
}
