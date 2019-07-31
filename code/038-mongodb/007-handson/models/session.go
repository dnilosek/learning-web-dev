package models

import "time"

type Session struct {
	UserName     string    `json:username`
	LastActivity time.Time `json:lastactivity`
}
