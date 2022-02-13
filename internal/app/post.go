package app

// Post ..
type Post struct {
	ID         int
	Votes      int
	Rate       int
	AuthorID   int
	Author     string
	Title      string
	Content    string
	Categories []Category
	Comments   []Comment
}
