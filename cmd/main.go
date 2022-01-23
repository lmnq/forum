package main

import (
	"forum/internal/service"
	"forum/internal/store"
	"log"
	"net/http"
)

func main() {
	db, err := store.InitDB()
	if err != nil {
		log.Println(err)
	}
	defer db.DB.Close()
	forum := service.Forum{
		DB: db,
	}

	http.HandleFunc("/", forum.IndexHandler)
	http.HandleFunc("/signin", forum.SignInHandler)
	fileServer := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/", http.StripPrefix("/static/css", fileServer))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
