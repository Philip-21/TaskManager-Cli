package database

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const DBPORT = 0600

var db *bolt.DB
var taskBucket = []byte("todo")

type Task struct {
	Key   int
	Value string
}

// int to bytes
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// bytes to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func ConnectDB(dbpath string) error {
	db, err := bolt.Open("task.db", DBPORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			return fmt.Errorf("Error in creating  bucket: %s", err)
		}
		return nil
	})
	return err

}

func Create(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id_64, _ := b.NextSequence() /// returns an autoincrementing integer for the bucket.
		id = int(id_64)
		key := itob(int(id_64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		cr := b.Cursor() //creates a cursor associated with the bucket   (cursor is to retrieve data, one row at a time, from a result set)
		for k, v := cr.First(); k != nil; k, v = cr.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k), //converts the bytes to an integer
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}
