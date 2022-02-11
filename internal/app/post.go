package app

// Post ..
type Post struct {
	ID         int
	Votes      int
	Status     int
	Author     string
	Title      string
	Content    string
	Categories []Category
}
