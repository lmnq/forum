package handlers

import (
	"fmt"
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"net/http"
)

// CreatePostGetHandler ..
func (f *Forum) CreatePostGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/createpost.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	categories, err := f.Service.GetAllCategories()
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, categories)
}

// CreatePostPostHandler ..
func (f *Forum) CreatePostPostHandler(ctx *router.Context) {
	tmpCategories, err := f.Service.GetCategoriesFromInput(ctx.Request.Form["category"])
	if err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	post := app.Post{
		Title:      ctx.Request.FormValue("title"),
		Content:    ctx.Request.FormValue("content"),
		Categories: tmpCategories,
	}
	categories, err := f.Service.GetAllCategories()
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	if err := f.Service.ValidatePostInput(post, categories); err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	postID, err := f.Service.AddNewPost(post)
	http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/post/%d", postID), 302)
}
