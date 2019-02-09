package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("hello world")

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Println(err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	items, err := fetchTopStories(ctx, 30)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := insertOrReplaceItems(db, items)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.RowsAffected())
}
