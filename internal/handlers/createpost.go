package handlers

import (
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
	post := &app.Post{
		Title:      ctx.Request.FormValue("title"),
		Content:    ctx.Request.FormValue("content"),
		Categories: ctx.Request.Form["category"],
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
	// addnewpost
}
