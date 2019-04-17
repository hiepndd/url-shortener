package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var urlshortenBucket = []byte("urlshorten")

var db *bolt.DB

// URLShorten is struct contain infor of specific url
type URLShorten struct {
	Key   string
	Value string
}

//Init is function create a DB
func Init(dbPath string) error {
	var err error
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(urlshortenBucket)
		return err
	})

}

// AddURLShorten is func add url to DB
func AddURLShorten(key, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(urlshortenBucket)
		return b.Put([]byte(key), []byte(value))
	})
	if err != nil {
		return err
	}
	return nil
}
