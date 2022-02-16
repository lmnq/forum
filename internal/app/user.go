package app

import "time"

// User ..
type User struct {
	Logged       bool
	ID           int
	Username     string
	Email        string
	Password     string
	HashPassword string
	Created      time.Time
}
