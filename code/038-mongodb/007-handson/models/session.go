package models

import "time"

type Session struct {
	UserID       string    `bson:userid`
	LastActivity time.Time `bson:lastactivity`
}
