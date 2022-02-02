package handlers

import (
	"forum/internal/app"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type postData struct {
	Post     *app.Post
	Comments []*app.Comment
}

// PostGetHandler ..
func (f *Forum) PostGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}
	post, err := f.Service.GetPost(id)
	if err != nil {
		log.Println(err)
		return
	}
	comments, err := f.Service.GetCommentsToPost(post)
	if err != nil {
		log.Println(err)
		return
	}
	data := &postData{
		Post:     post,
		Comments: comments,
	}
	tmpl.Execute(w, data)
	return
}
