package router

import (
	"html/template"
	"net/http"
)

// Context ..
type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

// WriteError ..
func (ctx *Context) WriteError(status int) {
	type ans struct {
		Code int
		Text string
	}
	w := ctx.ResponseWriter
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}
	switch status {
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		temp.Execute(w, ans{
			Code: status,
			Text: "Bad Request",
		})
	case 404:
		w.WriteHeader(http.StatusNotFound)
		temp.Execute(w, ans{
			Code: status,
			Text: "Page Not Found",
		})
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
		temp.Execute(w, ans{
			Code: status,
			Text: "Method Not Allowed",
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		temp.Execute(w, ans{
			Code: status,
			Text: "Internal Server Error",
		})
	}
}
