package indexes

import (
	"sync"

	"github.com/dylanlott/pangolin/pkg/collection"
	"github.com/dylanlott/pangolin/pkg/persist"
)

// Index holds the information for editing and updating indexes
// on a collection.
// Indexes have a RWMutex for safe manipulation.
type Index struct {
	Collection *collection.Collection
	Field      string
	Hashtable  ValueHashtable

	sync.RWMutex
}

// New creates an index on a given field and returns a pointer
// to that index.
func New(field string, col *collection.Collection) (*Index, error) {
	ht := ValueHashtable{}

	return &Index{
		Collection: col,
		Field:      field,
		Hashtable:  ht,
	}, nil
}

// Open returns a pointer to an Index struct for methods on indexes
func Open(field string, col *collection.Collection) (*Index, error) {
	var vh ValueHashtable

	if persist.Exists(col.Path) {
		// it exists, load the existing one
		err := persist.Load(col.Path, vh)
		if err != nil {
			return &Index{}, err
		}
	} else {
		// it doesn't exist, create a new one
		i, err := New(field, col)
		if err != nil {
			return nil, err
		}
		return i, nil
	}

	return &Index{
		Collection: col,
		Field:      field,
		Hashtable:  vh,
	}, nil
}

// Put inserts a value into the *Index
func (i *Index) Put(key string, value interface{}) (interface{}, error) {
	err := i.Hashtable.Put(key, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Get returns a value from a *Index
func (i *Index) Get(key string) (interface{}, error) {
	return i.Hashtable.Get(key)
}

// Delete removes a value from the Index
func (i *Index) Delete(key string) error {
	return i.Hashtable.Remove(key)
}
