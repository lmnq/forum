package router

import (
	"html/template"
	"net/http"
)

// Context ..
type Context struct {
	http.ResponseWriter
	*http.Request
	Params map[string]string
}

// setURLValues ..
func (ctx *Context) setURLValues(keys, values []string) {
	for i, key := range keys {
		ctx.SetParam(key, values[i])
	}
}

// SetParam ..
func (ctx *Context) SetParam(key, value string) {
	ctx.Params[key] = value
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
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
		temp.Execute(w, ans{
			Code: status,
			Text: "Unauthorized",
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		temp.Execute(w, ans{
			Code: status,
			Text: "Internal Server Error",
		})
	}
}
