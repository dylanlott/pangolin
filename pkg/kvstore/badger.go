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
func (k *KVStore) Get(key string) ([]byte, error) {
	var valCopy []byte
	err := k.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		// copy value, don't assign it
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, err
	}

	return valCopy, nil
}

// BulkGet takes as many keys as you pass it and returns their associated values.
// This acquires a lock before obtaining the keys.
func (k *KVStore) BulkGet(key ...string) ([][]byte, error) {
	panic("not implemented")
}

// Put puts a key value pair to disk
// This obtains a lock before putting values to disk.
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
