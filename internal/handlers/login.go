package handlers

import (
	"forum/internal/app"
	"forum/internal/router"
	"html/template"
	"log"
	"net/http"
)

// LoginGetHandler ..
func (f *Forum) LoginGetHandler(ctx *router.Context) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusInternalServerError)
		return
	}
	tmpl.Execute(ctx.ResponseWriter, nil)
}

// LoginPostHandler ..
func (f *Forum) LoginPostHandler(ctx *router.Context) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		return
	}
	user := app.User{
		Email:    ctx.Request.FormValue("email"),
		Password: ctx.Request.FormValue("password"),
	}
	// if err := f.Service.IsValidLoginData(user); err != nil {
	// 	log.Println(err)
	// 	ctx.WriteError(http.StatusBadRequest)
	// 	return
	// }
	if err := f.Service.LoginUser(user); err != nil {
		log.Println("fesf")
		log.Println(err)
		ctx.WriteError(http.StatusBadRequest)
		// 500
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
