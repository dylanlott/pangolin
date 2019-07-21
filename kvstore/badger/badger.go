package kvstore

import (
	"log"

	badger "github.com/dgraph-io/badger"
	"github.com/zeebo/errs"
)

// Pair fulfills KeyPair
type Pair struct {
	key   Keyed
	value Valued
}

// BadgerAdapter fulfills `KeyValueStore`
type BadgerAdapter struct {
	db *badger.DB
}

// NewBadgerAdapter Returns a new BadgerAdapter
func NewBadgerAdapter(path string) (*BadgerAdapter, error) {
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &BadgerAdapter{
		db: db,
	}, nil
}

// Get returns a KeyPair
func (b *BadgerAdapter) Get(key Keyed) (Paired, error) {
	return &Pair{}, nil
}

// Put inserts a KeyPair into the database, and returns the KeyPair
// or an error if it was unsuccessful.
func (b *BadgerAdapter) Put(key Keyed, value Valued) error {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	err := b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key.Bytes(), value.Bytes())
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

// Key returns the Keyed value or an error
func (p *Pair) Key() (Keyed, error) {
	panic("not impl")
}

// Value returns the Valued value or an error
func (p *Pair) Value() (Valued, error) {
	panic("not impl")
}
