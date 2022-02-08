package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
)

// IndexHandler ..
func (f *Forum) IndexHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	type data struct {
		Post     *app.Post
		Comments []*app.Comment
	}
	posts, err := f.Service.GetAllPosts(2)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, posts)
}
