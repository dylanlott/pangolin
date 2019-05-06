package kvstore

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/dgraph-io/badger"
)

// KVStore holds the active reference to the store and other details
type KVStore struct {
	DB   *badger.DB
	Path string
}

// Open returns a struct for interacting with Badger
func Open(path string) (*KVStore, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &KVStore{
		DB:   db,
		Path: path,
	}, err
}

// Get returns the value of a given Key.
func (k *KVStore) Get(key string) error {
	panic("not implemented")
}

// Put puts a key value pair to disk
func (k *KVStore) Put(key string, value interface{}) error {
	valBytes, err := getBytes(value)
	if err != nil {
		return err
	}
	keyBytes, err := getBytes(key)
	if err != nil {
		return err
	}
	return k.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(keyBytes, valBytes)
		return err
	})
}

// getBytes reads in a value and returns the byte value or an error
// This is necessary to store arbitrary interfaces in Badger
func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
