package apub

import (
	"time"
)

type Post interface {
	HTML() string
	Text() string
	Mentions() []string
	PostedAt() time.Time
	PostURL() string
	FeedURL() string
}
