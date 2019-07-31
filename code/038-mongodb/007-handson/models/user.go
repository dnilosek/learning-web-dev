package models

type User struct {
	UserName string `json:username`
	Password []byte `json:password`
	First    string `json:first`
	Last     string `json:last`
	Role     string `json:role`
}
