package main

import (
	"forum/internal/store"
	"log"
)

func main() {
	db, err := store.InitDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}
