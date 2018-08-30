package db

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/derekparker/trie"
	"github.com/zeebo/errs"
	homedir "github.com/mitchellh/go-homedir"
)

var (
	Error = errs.Class("pangolin_db_error")
	pangolinDB bytes.Buffer
	mutex = &sync.Mutex{}
	t trie.Trie
)

type Document struct {
	Id	string
	Data interface{}
}

type Response struct {
	Data    interface{}
	Success bool
	Meta		interface{}
}

type Collection struct {
	Name string
	Meta []Document
	trie trie.Trie
}

type DB struct {
	trie trie.Trie
	collections []*Collection
}

func NewCollection(name string) (*Collection) {
	return &Collection{
		Name: name,
		Meta: nil,
	}
}

func getPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".pangolin")
	return path, nil
}

func SaveCollection(name string, coll Collection) error {
	home, err := getPath() 
	path := filepath.Join(home, name)
	log.Printf("saving to collection %s at %s", name, path)
	err = Save(path, coll)
	if err != nil {
		return Error.New("Error saving collection")
	}
	return err
}

func LoadCollection(name string) (Collection) {
	dir, err := homedir.Dir()
	path := filepath.Join(dir, ".pangolin")
	var coll Collection
	err = Load(path, coll)
	if err != nil {
		log.Fatal("Error loading collection", err)
	}
	return coll
}

// NewDatabase will create a new database pointer
func NewDatabase() (*DB, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	t := trie.New()
	f := filepath.Join(dir, ".pangolin")

	log.Printf("starting new pangolin DB\n trie: %+v\n filepath: %+v\n", t, f)

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
	mutex.Lock()
	db.trie.Remove(key)
	db.trie.Add(key, data)
	mutex.Unlock()

	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// Delete will delete an object from the tree
func (db *DB) Delete(key string) (Response, error) {
	mutex.Lock()
	db.trie.Remove(key)
	ok := db.trie.HasKeysWithPrefix(key)
	mutex.Unlock()
	if !ok {
		log.Printf("deleted key %s", key)
		return Response{
				Data: key,
				Success: true,
		}, nil
	}
	log.Printf("could not delete key %s", key)

	return Response{
		Data: nil, 
		Success: false,
	}, Error.New("Delete Error: Key not deleted")
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

func (db *DB) createCollection(collection string) error {
	home, err := db.PangolinHomeDir()
	if err != nil {
		return err
	}
	collectionPath := filepath.Join(home, collection)
	createFile(collectionPath)
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
