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
	go store.DeleteExpiredSession(forumdb.DB)
	srv := service.NewService(forumdb)
	forum := handlers.NewForum(srv)
	defer forum.Service.Store.DB.Close()

	r := router.NewRouter()

	r.GET("/all", forum.IndexHandler)
	r.GET("/register", forum.RegisterGetHandler)
	r.POST("/register", forum.RegisterPostHandler)
	r.GET("/login", forum.LoginGetHandler)
	r.POST("/login", forum.LoginPostHandler)
	r.GET("/post/create", forum.CreatePostGetHandler)
	r.POST("/post/create", forum.AuthMiddleware(forum.CreatePostPostHandler))
	r.GET("/post/:postID", forum.PostGetHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.Handle("/", r)
	fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
