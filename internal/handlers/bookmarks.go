package handlers

import (
	"forum/internal/router"
	"html/template"
	"net/http"
	"strconv"
)

// BookmarksHandler ..
func (f *Forum) BookmarksHandler(ctx *router.Context) {
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	tmpl, err := template.ParseFiles("templates/bookmarks.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	posts, err := f.Service.GetBookmarkedPosts(userID)
	if err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, posts)
}