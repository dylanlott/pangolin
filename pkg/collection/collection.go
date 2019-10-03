package collection

import (
	"errors"
	"pangolin/pkg/indexes"
	"pangolin/pkg/kvstore"

	"github.com/zeebo/errs"
)

const (
	// ErrExist is returned if an object or item doesn't exist
	ErrExist = errs.Class("ErrExist")
)

// Collection holds the active Collection and its related data
type Collection struct {
	Name    string
	KV      kvstore.KeyValueStore
	Indexes map[string]*indexes.Index
}

// NewCollection returns a new *Collection
func NewCollection(name string, dir string) (*Collection, error) {
	kv, err := kvstore.BadgerKVStore(dir)
	return &Collection{
		Name: name,
		KV:   kv,
	}, err
}

// Get returns the value from Collection at Key `key` returns the Value
func (c *Collection) Get(key kvstore.Key) (kvstore.Value, error) {
	idx, ok := c.Indexes[c.Name]
	if !ok {
		return kvstore.Value{}, errors.New("index does not exist")
	}

	return idx.Get(key)
}
