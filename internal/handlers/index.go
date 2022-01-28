package handlers

import (
	"forum/internal/app"
	"html/template"
	"log"
	"net/http"
)

// IndexHandler ..
func (f *Forum) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	type data struct {
		Post     *app.Post
		Comments []*app.Comment
	}
	posts, err := f.Service.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, posts)
}
