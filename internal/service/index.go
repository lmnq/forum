package service

import (
	"fmt"
	"forum/internal/app"
	"forum/internal/store"
	"html/template"
	"log"
	"net/http"
)

// Forum ..
type Forum struct {
	DB *store.ForumDB
}

// IndexHandler ..
func (f *Forum) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/comments.html")
	if err != nil {
		log.Println(err)
		return
	}
	type data struct {
		Post     *app.Post
		Comments []*app.Comment
	}
	ans := []data{}
	// res := make(map[*app.Post][]*app.Comment)
	posts, err := f.DB.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(posts)
	for _, p := range posts {
		pComments, err := f.DB.GetCommentsToPost(p)
		// fmt.Println("comms: ", pComments[0].Content)
		// fmt.Println(p.Title)
		if err != nil {
			log.Println(err)
			return
		}
		// dodelat'
		// res[p] = pComments
		ans = append(ans, data{
			Post:     p,
			Comments: pComments,
		})
	}
	fmt.Println(ans)
	for _, d := range ans {
		fmt.Println(d.Post.Title)
		fmt.Println(d.Comments[0].Content)
	}
	// tmpl.Execute(w, res)
	tmpl.Execute(w, ans)
}
