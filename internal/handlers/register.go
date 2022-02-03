package handlers

import (
	"forum/internal/app"
	"html/template"
	"log"
	"net/http"
)

// RegisterGetHandler ..
func (f *Forum) RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
	return
}

// RegisterPostHandler ..
func (f *Forum) RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}
	user := &app.User {
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := f.Service.IsValidRegisterData(user); err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		log.Println(err)
		return
	}
	if err := f.Service.RegisterUser(user); err != nil {
		// error unique or 500
		// redirect to the same page with written data
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cookie, err := f.Service.SetCookie(user.Email)
	if err != nil {
		// Error
		log.Println(err)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/all", 301)
	return
}
