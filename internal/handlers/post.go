package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"strconv"
)

type postData struct {
	Post     *app.Post
	Comments []*app.Comment
}

// PostGetHandler ..
func (f *Forum) PostGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		return
	}
	// slug := router.GetField(r, 0)
	slug := ctx.Params[0]
	id, err := strconv.Atoi(slug)
	if err != nil {
		log.Println(err)
		return
	}
	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
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
	tmpl.Execute(ctx.ResponseWriter, data)
	return
}
