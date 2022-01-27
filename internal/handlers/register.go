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
	if r.Method == http.MethodPost {
		// Error
		if err := r.ParseForm(); err != nil {
			return
		}
		user, err := validateUserForm(r)
		if err != nil {
			log.Println(err)
			// Error
			return
		}
		err = f.Service.Store.RegisterUser(user)
		log.Println(err)
		http.Redirect(w, r, "/", 301)
		return
		// return
	}
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func validateUserForm(r *http.Request) (*app.User, error) {
	user := &app.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	return user, nil
}
