package kvstore

import (
	"log"

	badger "github.com/dgraph-io/badger"
	"github.com/zeebo/errs"
)

// BadgerAdapter fulfills `KeyValueStore`
type BadgerAdapter struct {
	db *badger.DB
}

var _ KeyValueStore = (*BadgerAdapter)(nil)

const (
	// ErrInvalidKey is returned if an empty key or invalid key is provided.
	ErrInvalidKey = errs.Class("invalid key")
)

// NewBadgerAdapter Returns a new BadgerAdapter
func NewBadgerAdapter(path string) (*BadgerAdapter, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return &BadgerAdapter{
		db: db,
	}, nil
}

// Get returns a KeyPair
func (b *BadgerAdapter) Get(key Key) (Value, error) {
	var value Value
	var err error
	b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		value = valCopy
		return nil
	})

	return value, err
}

// Put inserts a KeyPair into the database, and returns the KeyPair
// or an error if it was unsuccessful.
func (b *BadgerAdapter) Put(key Key, value Value) error {
	if key == nil {
		return errs.New("must provide non-nil key")
	}

	if value == nil {
		return errs.New("must provide value to insert")
	}

	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	err := b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// Delete removes a KeyPair from the KeyValueStore
func (b *BadgerAdapter) Delete(key Keyed) error {
	return errs.New("not impl")
}
