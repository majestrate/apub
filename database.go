package apub

// Database defines an interface for interacting with a persistant datastore for posts and user information
type Database interface {

	// LocalUser gets a local user by username
	// returns user, nil on found
	// returns nil, nil on not found
	// returns nil, error on db error
	LocalUser(username string) (*UserInfo, error)

	// LocalPost gets a local post by post id
	// returns post, nil on post found
	// returns nil, nil on post not found
	// returns nil, error on db error
	LocalPost(postid string) (*Post, error)

	// LocalUserPosts gets a slice of posts of a user by name with offset and limit
	// posts may be empty
	// retruns posts, nil on success
	// returns nil, err on db error
	LocalUserPosts(username string, offset int64, limit int) ([]*Post, error)

	ListFollowers(username string) ([]*UserInfo, error)
	ListFollowing(username string) ([]*UserInfo, error)
}
