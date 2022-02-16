package handlers

import (
	"forum/internal/router"
	"html/template"
	"net/http"
)

// BugHandler ..
func (f *Forum) BugHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/planet.html")
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, nil)
}
