package database

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const DBPORT = 0600

func ConnectDB() {
	db, err := bolt.Open("task.db", DBPORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	defer db.Close()
}

var db *bolt.DB

func Update() error {

	return nil
}

// func View() error {
// 	err = db.View(func(tx *bolt.Tx) error {
// 		conf := tx.Bucket()
// 	})
// }
