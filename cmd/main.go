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
	go store.CleanSessions(forumdb.DB)
	srv := service.NewService(forumdb)
	forum := handlers.NewForum(srv)
	defer forum.Service.Store.DB.Close()

	r := router.NewRouter()

	r.GET("/all", forum.AuthMiddleware(forum.IndexHandler, false))
	r.GET("/register", forum.RegisterGetHandler)
	r.POST("/register", forum.RegisterPostHandler)
	r.GET("/login", forum.LoginGetHandler)
	r.POST("/login", forum.LoginPostHandler)
	r.GET("/post/create", forum.CreatePostGetHandler)
	r.POST("/post/create", forum.AuthMiddleware(forum.CreatePostPostHandler, true))
	r.GET("/post/:postID", forum.AuthMiddleware(forum.PostGetHandler, false))
	r.POST("/post/:postID/comment", forum.AuthMiddleware(forum.CommentPostHandler, true))
	r.POST("/vote/:postID", forum.AuthMiddleware(forum.VotePostHandler, true))
	r.POST("/vote/:postID/:commentID", forum.AuthMiddleware(forum.VoteCommentHandler, true))
	r.GET("/category/:categoryID", forum.AuthMiddleware(forum.FilterByCategoryHandler, false))
	r.GET("/profile/:profileID/bookmarks", forum.AuthMiddleware(forum.BookmarksHandler, false))
	r.GET("/profile/:profileID/posts", forum.AuthMiddleware(forum.ProfilePostsHandler, false))
	r.GET("/bug", forum.BugHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.Handle("/", r)
	fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
