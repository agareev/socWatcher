package main

//
// import (
// 	"database/sql"
// )
//
// func push2Database(db *sql.DB, message string, id int) {
//   rows, err := db.Query(select "id from meet")
//   if err != nil {
//     log.Fatal(err)
//   }
//   defer rows.Close()
//   for rows.Next() {
//     err := rows.Scan(&id, &string)
//   }
// }

//   for {
// 		rows, err := db.Query(query)
// 		if err != nil {
// 			waitAndReconnect()
// 		}
// 		if rows != nil {
// 			state = true
// 		} else {
// 			state = false
// 		}
// 		time.Sleep(time.Second * 1)
// 	}
// }
