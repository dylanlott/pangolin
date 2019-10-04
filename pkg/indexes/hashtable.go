package indexes

import (
	"github.com/zeebo/errs"
	"golang.org/x/exp/mmap"
)

const (
	// ErrNoKey is thrown if an empty key is provided to the hashtable
	ErrNoKey = errs.Class("no key provided")
)

// Enforce HashTable to fulfill HT
var _ HT = (*HashTable)(nil)

// HT declares the interface all Hashtables need to
// adhere to.
type HT interface {
	Get(key string) (interface{}, error)
	Insert(key string, value interface{}) error
	Update(key string, value interface{}) error
	Delete(key string) error
}

// HashTable fulfills HT at runtime.
type HashTable struct {
	dir    string
	reader *mmap.ReaderAt
}

// NewHashTableIndex Returns a new HashTable or an error
func NewHashTableIndex(dir string) (*HashTable, error) {
	reader, err := mmap.Open(dir)
	if err != nil {
		return nil, errs.New("could not open mmap reader")
	}

	return &HashTable{
		dir:    dir,
		reader: reader,
	}, nil
}

// Get returns a key from the hash table
func (ht *HashTable) Get(key string) (interface{}, error) {
	return nil, errs.New("not impl")
}

// Insert puts a value into the hashtable
func (ht *HashTable) Insert(key string, value interface{}) error {
	return errs.New("not impl")
}

// Update finds a value and updates it. This should be immutable.
func (ht *HashTable) Update(key string, value interface{}) error {
	return errs.New("not impl")
}

// Delete removes a key value from the hashtable.
func (ht *HashTable) Delete(key string) error {
	return errs.New("not impl")
}
