package db

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"errors"

	"github.com/derekparker/trie"
	homedir "github.com/mitchellh/go-homedir"
)

// Response struct
type Response struct {
	Data    interface{}
	Success bool
}

// DB is the main struct exported out
type DB struct {
	trie trie.Trie
}

var pangolinDB bytes.Buffer
var mutex = &sync.Mutex{}
var t trie.Trie

// NewDatabase will create a new database pointer
func NewDatabase() (*DB, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	t := trie.New()
	f := filepath.Join(dir, ".pangolin")

	log.Printf("starting trie %+v\n in %+v\n", t, f)

	return &DB{
		trie: *t,
	}, nil
}

// Get returns object with id of `id`
func (db *DB) Get(key string) (Response, error) {
	mutex.Lock()
	node, ok := db.trie.Find(key)
	if !ok {
		fmt.Printf("error getting key %s", key)
	}
	mutex.Unlock()

	return Response{
		Data:    node.Meta(),
		Success: true,
	}, nil
}

// Insert puts a JSON blob into the collection
func (db *DB) Insert(key string, data interface{}) (Response, error) {
	db.trie.Add(key, data)
	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// Update inserts a value into the database
// If upsert is true, it will insert the data if the key is not found
func (db *DB) Update(key string, data interface{}) (Response, error) {
	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// Delete will delete an object from the tree
func (db *DB) Delete(key string) error {
	mutex.Lock()
	db.trie.Remove(key)
	node, ok := db.trie.Find(key)
	mutex.Unlock()
	if node == nil && ok {
		return nil
	}
	return errors.New("Delete Error: Key not deleted")
}

func (db *DB) getTrie() (trie.Trie, error) {
	return db.trie, nil
}

func (db *DB) PangolinHomeDir() (string, error) {
	dir, err := homedir.Dir()
	f := filepath.Join(dir, ".pangolin")
	return f, err
}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err) //okay to call os.exit()
		defer file.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
