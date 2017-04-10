package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func checkUniq(db *sql.DB, absNumber int) bool {
	rows, err := db.Query("SELECT abs_number FROM comments WHERE abs_number = ?", absNumber)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	if rows != nil {
		return true
	}
	return false
}

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

		// Проверка на uniq number в посте
		if checkUniq(db, absNumber) == false {
			res, err := stmt.Exec(absNumber, comment)
			if err != nil {
				log.Println(err)
			}
			id, err := res.LastInsertId()
			if err != nil {
				log.Println(err)
			}
			log.Println(id, "was wrote")
		}
	}
}
