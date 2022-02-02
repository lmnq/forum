package main

import (
	"fmt"
	"forum/internal/handlers"
	"forum/internal/router"
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

	r := router.NewRouter()
	r.GET("/register", forum.RegisterGetHandler)
	r.POST("/register", forum.RegisterPostHandler)
	r.GET("/login", forum.LoginGetHandler)
	// r.POST("/login", forum.LoginGetHandler)
	r.GET("/all", forum.IndexHandler)
	r.POST("/all", forum.IndexHandler)
	// http.HandleFunc("/", forum.IndexHandler)
	// http.HandleFunc("/login", forum.LoginHandler)
	// http.HandleFunc("/register", forum.RegisterHandler)
	// http.HandleFunc("/post", forum.PostHandler)
	fileServer := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/", http.StripPrefix("/static/css", fileServer))
	fmt.Println("server running on http://localhost:8081")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
