package collection

import (
	"fmt"
	"os"
	"sync"

	"github.com/dylanlott/pangolin/pkg/kvstore"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/zeebo/errs"
)

// Collection holds all of the collection information
// and necessary mutexes
type Collection struct {
	Name string

	// Location of collection files
	Path string

	// Map of field names in this collection to AVL trees
	// Indexes    []indexes.Index
	// HashTables map[string]hashtable.Hash

	Driver *kvstore.KVStore

	sync.Mutex
}

// NewCollection creates a new collection and instantiates its indices
func NewCollection(name string) (*Collection, error) {
	// TODO: Check that collection exists
	path, err := getCollectionPath(name)
	if err != nil {
		return nil, err
	}
	kv, err := kvstore.Open(path)
	if err != nil {
		return nil, err
	}

	// TODO: Create index files and load them on here.
	coll := &Collection{
		Name:   name,
		Driver: kv,
		Path:   path,
	}

	return coll, nil
}

// GetCollection returns a collection pointer or an error
func GetCollection(name string) (*Collection, error) {
	if pathExists(name) {
		path, err := getCollectionPath(name)
		if err != nil {
			return nil, err
		}

		kv, err := kvstore.Open(path)
		if err != nil {
			return nil, err
		}

		return &Collection{
			Name:   name,
			Driver: kv,
			Path:   path,
		}, nil
	}

	return NewCollection(name)
}

// RemoveCollection will delete all data pertaining to a collection.
// This is not un-doable. Don't expose this to end users without confirmation.
func RemoveCollection(name string) error {
	if pathExists(name) {
		_, err := getCollectionPath(name)
		if err != nil {
			return err
		}

		// TODO: Call remove collection function here.
	}
	return errs.New("not impl yet")
}

// Returns the path of the
// TODO: Make this take env configs rather than be hard coded
func getCollectionPath(name string) (string, error) {
	// TODO: This should come from a Config that's passed through
	// this service on startup.
	homePath, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/%s/%s", homePath, "pangolindb", name)
	fmt.Printf("PATH: %s", path)
	return path, nil
}

// pathExists checks if a given path exists.
func pathExists(name string) bool {
	path, err := getCollectionPath(name)
	if err != nil {
		fmt.Printf("error getting path in pathExists: %s", err.Error())
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
