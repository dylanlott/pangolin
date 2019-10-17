package indexes

import (
	"fmt"

	"github.com/zeebo/errs"
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
	Get(key string) interface{}
	Put(key string, value interface{}) error
	Update(key string, value interface{}) error
	Delete(key string) error
}

// HashTable fulfills HT at runtime.
type HashTable struct {
	items map[int]interface{}
	path  string
}

// NewHashTableIndex Returns a new HashTable or an error
func NewHashTableIndex(path string) (*HashTable, error) {
	ht := &HashTable{}
	ht.path = path
	err := Load(path, ht)
	if err != nil {
		return nil, err
	}

	return ht, err
}

// hashing algo for the table
func hash(k string) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

// Get returns a key from the hash table
func (ht *HashTable) Get(key string) interface{} {
	i := hash(key)
	return ht.items[i]
}

// Put puts a value into the hashtable
func (ht *HashTable) Put(key string, value interface{}) error {
	i := hash(key)
	ht.items[i] = value
	return Save(ht.path, ht)
}

// Update finds a value and updates it. This should be immutable.
func (ht *HashTable) Update(key string, value interface{}) error {
	return errs.New("not impl")
}

// Delete removes a key value from the hashtable.
func (ht *HashTable) Delete(key string) error {
	return errs.New("not impl")
}
