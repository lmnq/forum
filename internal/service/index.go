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
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	type data struct {
		post     *app.Post
		comments []*app.Comment
	}
	res := make(map[*app.Post][]*app.Comment)
	posts, err := f.DB.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(posts)
	for _, p := range posts {
		pComments, err := f.DB.GetCommentsToPost(p)
		fmt.Println("comms: ", pComments)
		fmt.Println(p.Title)
		if err != nil {
			log.Println(err)
			return
		}
		// dodelat'
		res[p] = pComments
		// res = append(res, data{
		// 	post:     p,
		// 	comments: pComments,
		// })
	}
	fmt.Println(res)
	tmpl.Execute(w, res)
}
