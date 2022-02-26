package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// PostData ..
type PostData struct {
	Post     app.Post
	Comments []app.Comment
}

// PostGetHandler ..
func (f *Forum) PostGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	postID, _ := strconv.Atoi(ctx.Params["postID"])
	post, err := f.Service.GetPost(postID)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusNotFound)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, post)
}
