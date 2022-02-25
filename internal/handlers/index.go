package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// IndexData ..
type IndexData struct {
	User  app.User
	Posts []app.Post
}

// IndexHandler ..
func (f *Forum) IndexHandler(ctx *router.Context) {
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	posts, err := f.Service.GetAllPosts(userID)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	data := IndexData{
		User: app.User{
			Logged: true,
			ID:     userID,
		},
		Posts: posts,
	}
	tmpl.Execute(ctx.ResponseWriter, data)
}
