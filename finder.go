package apub

// InfoFinder finds a UserInfo given a string
type InfoFinder interface {
	// returns nil, nil on not found
	// returns url, nil on found
	// returns nil, error on error
	LocalUser(string) (*UserInfo, error)

	// list all users that follow a user by local username
	ListFollowers(string) ([]*UserInfo, error)

	// list all users that a local users follows by local username
	ListFollowing(string) ([]*UserInfo, error)
}
