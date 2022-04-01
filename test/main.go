package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	bu := "wid"
	key := "mes"
	value := "trues"
	//DBStats()
	//CreateBucket(bu, key, "this")
	GetAllkv(bu)
	PutDB(bu, key, value)
	DeleteDB(bu, key)

}

// tempfile returns a temporary file path.
func tempfile() string {
	if err := os.MkdirAll("db", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}
	f, err := os.Create("db/bolt.db")
	if err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}

func CreateBucket(bu string, key string, value string) {
	fmt.Println("creating buicket...")
	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bu))
		return err
	}); err != nil {
		log.Fatal(err)
	}
	// Create several keys in a transaction.
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	b := tx.Bucket([]byte(bu))

	if err = b.Put([]byte(key), []byte(value)); err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// Iterate over the values in sorted key order.
	tx, err = db.Begin(false)
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Rollback(); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}

func GetAllkv(bu string) {
	fmt.Println("getting all keys...")

	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Open the database.

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bu))
		// we need cursor for iteration
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s likes %s\n", k, v)
		}
		// should return nil to complete the transaction
		return nil
	})

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}

func DBStats() {
	fmt.Println("running stats...")

	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	prev := db.Stats()

	// Wait for 10s.

	time.Sleep(10 * time.Second)

	// Grab the current stats and diff them.

	stats := db.Stats()

	diff := stats.Sub(&prev)

	// Encode stats to JSON and print to STDERR.

	json.NewEncoder(os.Stderr).Encode(diff)

	// Save stats for the next loop.

	prev = stats

}

func PutDB(bu string, key string, values string) {
	fmt.Println("put...")

	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b := tx.Bucket([]byte(bu))

		// Set the value "bar" for the key "foo".
		if err := b.Put([]byte(key), []byte(values)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Read value back in a different read-only transaction.
	if err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte(bu)).Get([]byte(key))
		fmt.Printf("The value of 'foo' is: %s\n", value)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Close database to release file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
func DeleteDB(bu string, key string) {
	fmt.Println("delete...")

	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Delete the key in a different write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bu)).Delete([]byte(key))
	}); err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bu))
		// we need cursor for iteration
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s likes %s\n", k, v)

		}
		// should return nil to complete the transaction
		return nil
	})

	// Close database to release the file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

}
