package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func push2Database(comments map[int]string) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	for absNumber, comment := range comments {
		stmt, err := db.Prepare("INSERT INTO comments(abs_number, comment) values(?,?)")
		if err != nil {
			log.Fatal(err)
		}

		// UNIQUE constraint failed: comments.abs_number
		// сделать проверку на уникальность absNumber
		res, err := stmt.Exec(absNumber, comment)
		if err != nil {
			log.Fatal(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}
