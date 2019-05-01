package collection

import (
	"fmt"
	"sync"

	"github.com/ancientlore/go-avltree"
	"github.com/dylanlott/pangolin/pkg/kvstore"
	"github.com/timtadh/data-structures/hashtable"
)

// Collection holds all of the collection information
// and necessary mutexes
type Collection struct {
	Name string

	// Map of field names in this collection to AVL trees
	Indexes    map[string]avltree.PairTree
	HashTables map[string]hashtable.Hash

	Driver kvstore.BadgerKVStore

	sync.Mutex
}

// NewCollection creates a new collection and instantiates its indices
func NewCollection(name string) (*Collection, error) {
	// TODO: Check that collection exists
	path := getCollectionPath(name)

	fmt.Printf("Getting Collection at %s", path)

	kv := badger.NewBadgerKVStore(path)

	coll := &Collection{
		Name:       name,
		Indexes:    make(map[string]avltree.PairTree),
		HashTables: make(map[string]hashtable.Hash),
		Driver:     kv,
	}

	return coll
}

// GetCollection returns a collection pointer or an error
func GetCollection(name string) (*Collection, error) {
	path := getCollectionPath(name)

	kv := badger.NewBadgerKVStore(path)

	// TODO: Load indexes and hashtables onto Collection

	return &Collection{
		Name:   name,
		Driver: kv,
		// TODO: Load indexes and hashtables onto here.
	}
}

func getCollectionPath(name string) string {
	return fmt.Sprintf("/tmp/pangolin/", name)
}

func collExists(name string) bool {
	panic("not impl")
}
