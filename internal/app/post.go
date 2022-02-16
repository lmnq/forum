package app

import "time"

// Post ..
type Post struct {
	ID         int
	Votes      int
	Rate       int
	AuthorID   int
	Author     string
	Title      string
	Content    string
	Created    time.Time
	Categories []Category
	Comments   []Comment
}
