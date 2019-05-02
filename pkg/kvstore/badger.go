package kvstore

import (
	"log"

	"github.com/dgraph-io/badger"
)

// BadgerKVStore holds the active reference to the store and other details
type BadgerKVStore struct {
	DB   *badger.DB
	Path string
}

// Open returns a struct for interacting with Badger
func Open(path string) (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &BadgerKVStore{
		DB:   db,
		Path: path,
	}, err
}

// Get creates a badger transaction and tries to lookup the value at the given key.
// If it doesn't exist, ok will be false. If it does exist, ok will be true, but the value
// could still be nil.
func (b *BadgerKVStore) Get(key string) (interface{}, bool, error) {
	var value interface{}
	var ok bool

	err := b.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy()
		if err != nil {
			return err
		}

		ok = true
		value = valCopy
		return nil
	})
	if err != nil {
		return nil, ok, err
	}
	return value, ok, err
}

// Put puts a key value pair to disk
func (b *BadgerKVStore) Put(key string, value interface{}) error {
	return b.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(string), []byte(value))
		return err
	})
}
