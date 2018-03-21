package apub

import "net/url"

// URLFinder finds a url given a string
// returns nil, nil on not found
// returns url, nil on found
// returns nil, error on error
type URLFinder func(string) (*url.URL, error)
