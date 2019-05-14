package hashtable

import (
	"github.com/dylanlott/pangolin/pkg/pangolin"
)

// SIZE is the number of buckets in the hashtable
const SIZE int = 15

// HashTable holds the table and size references.
type HashTable struct {
	Table map[string]*Node
	Size  int
}

// hash returns the hash of the key
func hash(i, size int) int {
	return (i % size)
}

// Node is a node in the hashtable
type Node struct {
	Value interface{}
	Next  *Node
}

// New returns a new Hashtable
func New() (*HashTable, error) {
	panic("not impl")
}

// Get returns the value at key Key in the hashtable.
// This acquires a lock on the hashtable.
func Get(key string) (*pangolin.Pair, error) {
	panic("not impl")
}

// Put puts a key / value pair to the hashtable.
// This acquires a lock before getting it.
func Put(key string, value interface{}) (interface{}, error) {
	panic("not impl")
}

// Size returns the size of the hashtable in an int64
func Size() int64 {
	panic("not impl")
}
