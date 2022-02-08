package handlers

import (
	"forum/internal/router"
	"html/template"
	"net/http"
)

// CreataPostGethandler ..
func (f *Forum) CreataPostGethandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/createpost.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, nil)
}
