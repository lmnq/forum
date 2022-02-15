package store

import "forum/internal/app"

// GetBookmarkedPosts ..
func (db *ForumDB) GetBookmarkedPosts(userID int) ([]app.Post, error) {
	posts := []app.Post{}

	rows, err := db.DB.Query(`
			SELECT
					posts.ID, posts.title, posts.content,
					users.username
			FROM
				post_votes AS pv INNER JOIN posts ON pv.post_ID = posts.ID
				INNER JOIN users ON posts.author_ID = users.ID
			WHERE
				pv.user_ID = ? AND pv.rate = 1
			ORDER BY pv.ID DESC;
	`, userID)

	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		categories, err := db.GetCategoriesToPost(post.ID)
		if err != nil {
			return posts, err
		}
		post.Categories = categories
		votes, rate, err := db.GetVotesToPost(post.ID, userID)
		if err != nil {
			return posts, err
		}
		post.Votes = votes
		post.Rate = rate
		posts = append(posts, post)
	}

	return posts, nil
}

// GetProfilePosts ..
func (db *ForumDB) GetProfilePosts(profileID int) ([]app.Post, error) {
	posts := []app.Post{}

	rows, err := db.DB.Query(`
			SELECT
					posts.ID, posts.title, posts.content,
					users.username
			FROM
				posts INNER JOIN users ON posts.author_ID = users.ID
			WHERE
				posts.author_ID = ?
			ORDER BY posts.ID DESC;
	`, profileID)

	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := app.Post{}
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
		categories, err := db.GetCategoriesToPost(post.ID)
		if err != nil {
			return posts, err
		}
		post.Categories = categories
		votes, rate, err := db.GetVotesToPost(post.ID, profileID)
		if err != nil {
			return posts, err
		}
		post.Votes = votes
		post.Rate = rate
		posts = append(posts, post)
	}

	return posts, nil
}
