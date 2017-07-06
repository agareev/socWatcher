package main

import (
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

func writeAllToDb(comments map[int]string) {
	for absNumber, comment := range comments {
		writeToDb(absNumber, comment)
	}
	log.Println("AllWrite!")
}

// Put comment
func Put(bucket, key string, val []byte) error {
	return dbc.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		if err = b.Put([]byte(key), val); err != nil {
			return err
		}
		return err
	})
}

// Get comment
func Get(bucket, key string) (data []byte, err error) {
	dbc.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		r := b.Get([]byte(key))
		if r != nil {
			data = make([]byte, len(r))
			copy(data, r)
		}
		return nil
	})
	return
}

func writeToDb(absNumber int, comment string) {

	err := Put("bucket", strconv.Itoa(absNumber), []byte(comment))
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("%+v\n", absNumber)

	data, err := Get("bucket", strconv.Itoa(absNumber))
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println(string(data))

}
