package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"net/http"
	"strconv"
)

// FilterData ..
type FilterData struct {
	Category app.Category
	Posts    []app.Post
}

// FilterByCategoryHandler ..
func (f *Forum) FilterByCategoryHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/category.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}

	categoryID, _ := strconv.Atoi(ctx.Params["categoryID"])
	category, err := f.Service.GetCategoryByID(categoryID)
	if err != nil {
		ctx.WriteError(http.StatusNotFound)
		return
	}

	userID, _ := strconv.Atoi(ctx.Params["userID"])
	posts, err := f.Service.GetPostsByCategory(categoryID, userID)
	if err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	data := FilterData{
		Category: category,
		Posts:    posts,
	}
	tmpl.Execute(ctx.ResponseWriter, data)
}
