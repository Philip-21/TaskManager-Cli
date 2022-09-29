package database 

import (
	"log"

	"github.com/boltdb/bolt"
)

const DBPORT = 0600

func ConnectDB() {
	db, err := bolt.Open("task.db", DBPORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
