package service

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
			// Error
			return
		}
		f.DB.RegisterUser(user)
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
		Username: r.FormValue("Username"),
		Email:    r.FormValue("Email"),
		Password: r.FormValue("Password"),
	}
	return user, nil
}
