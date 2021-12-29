package main

import (
	"fmt"
	"forum/internal/store"
	"log"
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
}

func printPosts(intr store.DataBase) {
	posts, err := intr.GetAllPosts()
	if err != nil {
		log.Println(err)
		return
	}
	for _, post := range posts {
		fmt.Printf("ID: %v\nTitle: %v\nContent: %v\nAuthorID: %v\n\n", post.ID, post.Title, post.Content, post.AuthorID)
	}
}
