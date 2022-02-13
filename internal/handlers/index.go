package handlers

import (
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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
	tmpl.Execute(ctx.ResponseWriter, posts)
}
