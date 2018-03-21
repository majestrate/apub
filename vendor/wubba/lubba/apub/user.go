package apub

type UserInfo interface {
	Subject() string
	Alias() string
	Links() []Link
}
