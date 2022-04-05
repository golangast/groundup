package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

func Tempfile() string {
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
func JustBucketCreator(bu string) {
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
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
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

func GetAllkv(bu string) ([]string, []string) {
	fmt.Println("getting all keys...")
	var titles []string
	var urls []string

	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(bu)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bu))
		// we need cursor for iteration
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s likes %s\n", k, v)
			titles = append(titles, string(k))
			urls = append(urls, string(v))

		}
		// should return nil to complete the transaction
		return nil
	})

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

	return titles, urls
}

func DBStats() {
	fmt.Println("running stats...")
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	prev := db.Stats()
	time.Sleep(10 * time.Second)
	stats := db.Stats()
	diff := stats.Sub(&prev)
	json.NewEncoder(os.Stderr).Encode(diff)
	prev = stats
}

func PutDB(bu string, key string, values string) *bolt.DB {
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

	return db
}

func AddDB(bu string, key string, value string) {
	fmt.Println("add...")

	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute several commands within a read-write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bu))

		if err := b.Put([]byte(key), []byte(value)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Read the value back from a separate read-only transaction.
	if err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte(bu)).Get([]byte(value))
		fmt.Printf("this value is saved: %s\n", value)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Close database to release the file lock.
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

func EmbeddBucket(originalbucket string, newbucket string) {
	fmt.Println("embedding...")
	// Open the database.
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Start the transaction.
	tx, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()
	root := tx.Bucket([]byte(originalbucket))
	// Setup the users bucket.

	if err := db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		bkt, err := root.CreateBucketIfNotExists([]byte(newbucket))
		if err != nil {
			return err
		}

		// Set the value "bar" for the key "foo".
		if err := bkt.Put([]byte("jim"), []byte("jannny")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Read value back in a different read-only transaction.
	if err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte(newbucket)).Get([]byte("jim"))
		fmt.Printf("The value of 'foo' is: %s\n", value)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	// Close database to release the file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

}
func CreateNestedBucketsNew(buckets []string) (err error) {
	db, err := bolt.Open("db/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) (err error) {
		var bkt *bolt.Bucket

		for index, bucket := range buckets {
			if index == 0 {
				bkt, err = tx.CreateBucketIfNotExists([]byte(bucket))
			} else {
				bkt, err = bkt.CreateBucketIfNotExists([]byte(bucket))
			}

			if err != nil {
				return fmt.Errorf("Error creating nested bucket [%s]: %v", bucket, err)
			}
		}
		return err
	})
	return err
}
