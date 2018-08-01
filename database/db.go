package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/derekparker/trie"
	homedir "github.com/mitchellh/go-homedir"
)

// Response struct
type Response struct {
	Data    interface{}
	Success bool
}

type DB struct {
	trie trie.Trie
}

var pangolinDB bytes.Buffer

// declare mutex for safe writes
var mutex = &sync.Mutex{}

var t trie.Trie

// NewDatabase will create a new database pointer
func NewDatabase() (*DB, error) {

	dir, err := homedir.Dir()
	if err != nil {
		log.Print("error finding home directory")
	}
	fmt.Printf(dir)

	if _, err := os.Stat(filepath.Join(dir, ".pangolin/pangolin.db")); err == nil {
		fmt.Printf("file exists")
		// load in file
		file, err := os.Open(filepath.Join(dir, ".pangolin/pangollin.db"))
		fmt.Printf("file %+v\n", file)
		fmt.Printf("error opening file", err)

		dec := gob.NewDecoder(&pangolinDB)
		err = dec.Decode(file)
		fmt.Printf("%q", &pangolinDB)
	} else {
		_, err := os.Create(filepath.Join(dir, ".pangolin/pangolin.db"))
		if err != nil {
			fmt.Printf("error creating file %s", err)
		}
	}

	//enc := gob.NewEncoder(&pangolinDB)

	t := trie.New()

	return &DB{
		trie: *t,
	}, nil
}

// Get returns object with id of `id`
func (db *DB) Get(key string) (Response, error) {
	node, ok := db.trie.Find(key)
	if !ok {
		fmt.Printf("error getting key %s", key)
	}
	meta := node.Meta()
	fmt.Printf("META::: %+v\n", meta)

	return Response{
		Data:    meta,
		Success: true,
	}, nil
}

// Query will find documents that match the query and return them
func (db *DB) Query() error {
	return nil
}

// Insert puts a JSON blob into the collection
func (db *DB) Insert(key string, data interface{}) (Response, error) {
	db.trie.Add(key, data)
	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// Put inserts a value into the database
// If upsert is true, it will insert the data if the key is not found
func (db *DB) Update(key string, data interface{}) (Response, error) {
	fmt.Printf("%+v\n", data)
	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// Delete will delete an object from the tree
func (db *DB) Delete(id string) error {
	return nil
}

func (db *DB) getTrie() (trie.Trie, error) {
	return db.trie, nil
}
