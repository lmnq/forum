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
	user := &app.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	err := f.Service.RegisterUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	cookie, err := f.Service.SetCookie(user.Email)
	if err != nil {
		// Error
		log.Println(err)
		return
	}
	http.SetCookie(w, cookie)
	// set cookie into db. delete if exists
	http.Redirect(w, r, "/all", 301)
	return
}
