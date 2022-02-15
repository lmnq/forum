package handlers

import (
	"fmt"
	"forum/internal/router"
	"html/template"
	"net/http"
	"strconv"
)

// FilterByCategoryHandler ..
func (f *Forum) FilterByCategoryHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/category.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	categoryID, _ := strconv.Atoi(ctx.Params["categoryID"])
	fmt.Println("categoryID: ", categoryID)
	userID, _ := strconv.Atoi(ctx.Params["userID"])

	posts, err := f.Service.GetPostsByCategory(categoryID, userID)
	if err != nil {
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, posts)
}
