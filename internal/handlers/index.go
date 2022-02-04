package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
)

// IndexHandler ..
func (f *Forum) IndexHandler(ctx *router.Context) {
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
	tmpl.Execute(ctx.ResponseWriter, posts)
}
