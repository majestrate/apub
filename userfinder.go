package apub

import "wubba/lubba/apub"

type Link = apub.Link
type UserInfo = apub.UserInfo

// InfoFinder finds a UserInfo given a string
// returns nil, nil on not found
// returns url, nil on found
// returns nil, error on error
type InfoFinder func(string) (UserInfo, error)
