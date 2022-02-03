package handlers

import (
	"forum/internal/app"
	"html/template"
	"log"
	"net/http"
)

// LoginGetHandler ..
func (f *Forum) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

// LoginPostHandler ..
func (f *Forum) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	user := &app.User{
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Service.IsValidLoginData(user); err != nil {
		log.Println(err)
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if err := f.Service.LoginUser(user); err != nil {
		log.Println(err)
		return
	}
	cookie, err := f.Service.SetCookie(user.Email)
	if err != nil {
		log.Println(err)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/all", 301)
}
