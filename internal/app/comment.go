package app

// Comment ..
type Comment struct {
	ID       int
	PostID   int
	Votes    int
	Rate     int
	AuthorID int
	Author   string
	Content  string
}
