package collection

import (
	"fmt"
	"os"
	"sync"

	"github.com/ancientlore/go-avltree"
	"github.com/dylanlott/pangolin/pkg/kvstore"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/timtadh/data-structures/hashtable"
	"github.com/zeebo/errs"
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
	kv := kvstore.NewBadgerKVStore(path)

	// TODO: Create index files and load them on here.
	coll := &Collection{
		Name:   name,
		Driver: kv,
	}

	return coll
}

// GetCollection returns a collection pointer or an error
func GetCollection(name string) (*Collection, error) {
	path := getCollectionPath(name)

	kv := kvstore.NewBadgerKVStore(path)

	// TODO: Load indexes and hashtables onto Collection

	return &Collection{
		Name:   name,
		Driver: kv,
	}
}

// RemoveCollection will delete all data pertaining to a collection.
// This is not un-doable. Don't expose this to end users without confirmation.
func RemoveCollection(name string) error {
	path := getCollectionPath(name)
	return errs.New("not implemented yet")
}

// Returns the path of the
// TODO: Make this take env configs rather than be hard coded
func getCollectionPath(name string) string {
	path := fmt.Sprintf("%s/%s/%s", homedir.Dir(), "pangolindb", name)
	fmt.Printf("PATH: %s", path)
	return path
}

// pathExists checks if a given path exists.
func pathExists(name string) bool {
	path := getCollectionPath(name)
	fmt.Printf("Checking path: %s", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
