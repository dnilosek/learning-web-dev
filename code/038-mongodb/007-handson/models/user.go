package models

type User struct {
	UserName string `bson:username`
	Password []byte `bson:password`
	First    string `bson:first`
	Last     string `bson:last`
	Role     string `bson:role`
}
