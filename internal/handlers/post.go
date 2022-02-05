package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
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
func (f *Forum) PostGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		return
	}
	idparam := ctx.Params[0]
	id, _ := strconv.Atoi(idparam)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	post, err := f.Service.GetPost(id)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusNotFound)
		return
	}
	comments, err := f.Service.GetCommentsToPost(post)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	data := &postData{
		Post:     post,
		Comments: comments,
	}
	tmpl.Execute(ctx.ResponseWriter, data)
}
