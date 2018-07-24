package db

import (
	"fmt"

	"github.com/derekparker/trie"
)

// Response struct
type Response struct {
	Data    interface{}
	Success bool
}

type DB struct {
	trie trie.Trie
}

// NewDatabase will create a new database pointer
func NewDatabase() (*DB, error) {
	t := trie.New()
	fmt.Printf("trie: %+v\n", t)
	return &DB{
		trie: *t,
	}, nil
}

// Get returns object with id of `id`
func (db *DB) Get(key string) (Response, error) {
	node, ok := db.trie.Find(key)
	fmt.Printf("ok is %+v\n", ok)
	fmt.Printf("found node %+v\n", node)
	if !ok {
		fmt.Printf("error getting key %s", key)
	}
	meta := node.Meta()
	fmt.Printf("META::: %+v\n", meta)

	return Response{
		Data:    node,
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
func (db *DB) Delete(id int) error {
	return nil
}

func (db *DB) getTrie() (trie.Trie, error) {
	return db.trie, nil
}
