package user

import (
	"net/url"
)

type FinderFunc func(string) (*url.URL, error)
