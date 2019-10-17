package indexes

/*
* Indexes TODO:
* - Make indexes mmap backed
* - Introduce transaction management
 */

import (
	"pangolin/pkg/kvstore"
)

// Index holds the information about the current index
type Index struct {
	dir string
}

// BuildIndex returns an index from the provided data.
func BuildIndex() {

}

// GetIndex returns an *Index
func GetIndex() {

}

// NewIndex returns a new Index
func NewIndex() *Index {
	return &Index{}
}

// Insert adds an item to an *Index
func (i *Index) Insert() {}

// Remove deletes an item from the index
func (i *Index) Remove() {}

// Get returns the value of the given key or an error
func (i *Index) Get(key kvstore.Key) (kvstore.Value, error) {
	return kvstore.Value{}, nil
}
