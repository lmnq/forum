package handlers

import (
	"fmt"
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	userID, _ := strconv.Atoi(ctx.Params["userID"])
	ctx.Request.ParseForm()
	tmpCategories, err := f.Service.GetCategoriesFromInput(ctx.Request.Form["category"])
	if err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	post := app.Post{
		AuthorID:   userID,
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
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	postID, err := f.Service.AddNewPost(post)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	// ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/post/%d", postID), 302)
}
