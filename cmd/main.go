package main

import (
	"fmt"
	"forum/internal/service"
	"forum/internal/store"
	"log"
	"net/http"
	"reflect"
)

func main() {
	db, err := store.InitDB()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reflect.TypeOf(db))
	// fmt.Println(db.GetAllPosts())
	printPosts(db)
	defer db.DB.Close()
	forum := service.Forum{
		DB: db,
	}

	http.HandleFunc("/", forum.IndexHandler)
	// http.HandleFunc("/", service.IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printPosts(intr store.DataBase) {
	posts, err := intr.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	for _, post := range posts {
		fmt.Printf("ID: %v\nTitle: %v\nContent: %v\nAuthorID: %v\n\n", post.ID, post.Title, post.Content, post.Author)
	}

}
