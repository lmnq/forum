package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
)

// RegisterGetHandler ..
func (f *Forum) RegisterGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, nil)
	return
}

// RegisterPostHandler ..
func (f *Forum) RegisterPostHandler(ctx *router.Context) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	user := &app.User{
		Username: ctx.Request.FormValue("username"),
		Email:    ctx.Request.FormValue("email"),
		Password: ctx.Request.FormValue("password"),
	}
	if err := f.Service.IsValidRegisterData(user); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	if err := f.Service.RegisterUser(user); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	cookie, err := f.Service.SetCookie(user.Email)
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	http.SetCookie(ctx.ResponseWriter, cookie)
	http.Redirect(ctx.ResponseWriter, ctx.Request, "/all", 302)
}
