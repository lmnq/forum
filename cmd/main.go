package main

import (
	"forum/internal/handlers"
	"forum/internal/service"
	"forum/internal/store"
	"log"
	"net/http"
)

func main() {
	forumdb, err := store.NewDataBase()
	if err != nil {
		log.Fatal(err)
	}
	srv := service.NewService(forumdb)
	forum := handlers.NewForum(srv)
	defer forum.Service.Store.DB.Close()

	http.HandleFunc("/", forum.IndexHandler)
	http.HandleFunc("/login", forum.LoginHandler)
	http.HandleFunc("/register", forum.RegisterHandler)
	fileServer := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/", http.StripPrefix("/static/css", fileServer))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
