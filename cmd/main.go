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

	// r := router.NewRouter()
	// r.GET("/register", forum.RegisterGetHandler)
	// r.POST("/register", forum.RegisterPostHandler)
	// r.GET("/login", forum.LoginGetHandler)
	// // r.POST("/login", forum.LoginGetHandler) ..
	// r.GET("/post", forum.PostGetHandler)
	// r.GET("/all", forum.IndexHandler)
	// r.POST("/all", forum.IndexHandler)
	// http.HandleFunc("/", forum.IndexHandler)
	// http.HandleFunc("/login", forum.LoginHandler)
	// http.HandleFunc("/register", forum.RegisterHandler)
	// http.HandleFunc("/post", forum.PostHandler)

	// regex table
	// re := router.NewReRouter()
	// re.NewRoute("GET", "/all", forum.IndexHandler)
	// re.NewRoute("GET", "/post/([0-9]+)", forum.PostGetHandler)
	// re.NewRoute("GET", "/login", forum.LoginGetHandler)
	// http.Handle("/", re)

	cr := router.NewRouter()
	cr.GET("/all", forum.IndexHandler)
	cr.GET("/post/([0-9]+)", forum.PostGetHandler)
	http.Handle("/", cr)

	fileServer := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/", http.StripPrefix("/static/css", fileServer))
	fmt.Println("server running on http://localhost:8080")
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
