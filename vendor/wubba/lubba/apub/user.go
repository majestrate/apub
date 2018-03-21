package apub

// serializable user info
type UserInfo interface {
	Subject() string
	Alias() string
	Links() []Link
}
