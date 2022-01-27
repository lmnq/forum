package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// LoginHandler ..
func (f *Forum) LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Println(err)
		return
	}
	// // Creating UUID Version 4
	// // panic on error
	// u1 := uuid.Must(uuid.NewV4(), nil)
	// fmt.Printf("UUIDv4: %s\n", u1)

	// // or error handling
	// u2 := uuid.NewV4()
	// if err != nil {
	// 	fmt.Printf("Something went wrong: %s", err)
	// 	return
	// }
	// fmt.Printf("UUIDv4: %s\n", u2)

	// // Parsing UUID from string input
	// u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// if err != nil {
	// 	fmt.Printf("Something went wrong: %s", err)
	// 	return
	// }
	// fmt.Printf("Successfully parsed: %s", u2)
	tmpl.Execute(w, nil)
}
