package main

import (
	"forum/internal/handlers"
	"log"
	"net/http"
)

func main() {
	forum, err := handlers.NewForum()
	if err != nil {
		log.Println(err)
		return
	}
	defer forum.Service.Store.DB.Close()

	http.HandleFunc("/", forum.IndexHandler)
	http.HandleFunc("/login", forum.LoginHandler)
	http.HandleFunc("/register", forum.RegisterHandler)
	fileServer := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/", http.StripPrefix("/static/css", fileServer))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
