package store

import "forum/internal/app"

// GetAllPosts ..
func (db *ForumDB) GetAllPosts(userID int) ([]*app.Post, error) {
	posts := []*app.Post{}
	rows, err := db.DB.Query(`
	SELECT
			posts.ID,
			title,
			content,
			username
	FROM posts INNER JOIN users ON posts.authorID = users.ID;
	`)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := &app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		// categories, err := db.GetCategoriesToPost(post.ID)
		// if err != nil {
		// 	return posts, err
		// }
		// post.Categories = categories
		votes, err := db.GetVotesToEntity("post", post.ID, userID)
		if err != nil {
			return posts, err
		}
		post.Votes = votes[0]
		post.Status = votes[1]
		posts = append(posts, post)
	}
	return posts, nil
}