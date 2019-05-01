package indexes

import (
	"github.com/timtadh/data-structures/hashtable"
)

// Client is the main interface that packages need to support
// to index documents
type Client interface {
	Get(key Key) (Pair, error)
	Put(key Key, value Value) error
	Remove(key string) error
	Size(key string) (int64, error)
	Load(name string) (*Index, error)
}

// Index tracks a specific index on a field of a collection.
type Index struct {
	collection *Collection
	field      string
}

// NewIndex creates a new index and returns a pointer or an error to that index
func NewIndex() *Index {
	return hashtable.NewHashTable(16)
}

// LoadIndex loads the index and returns a pointer to it or an error
func LoadIndex(collection string, field string) (*Index, error) {
}

// Get atomically reads an item from an Index
func (i *Index) Get() {

}

// Put atomically inserts or updates a key/value pair.
func (i *Index) Put() {

}

// Delete atomically removes an item from an index
func (i *Index) Delete() {

}
