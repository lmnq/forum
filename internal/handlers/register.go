package handlers

import (
	"forum/internal/app"
	"html/template"
	"log"
	"net/http"
)

// RegisterHandler ..
func (f *Forum) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		// Error
		return
	}
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	if r.Method != http.MethodPost {
		// Error
		w.WriteHeader(405)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		return
	}
	user := &app.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	err := f.Service.Store.RegisterUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", 302)
	return
}
