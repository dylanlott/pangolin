package db

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/derekparker/trie"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/zeebo/errs"
)

var (
	// Error is the main error class for this package
	Error      = errs.Class("ERRPANGOLIN")
	pangolinDB bytes.Buffer
	mutex      = &sync.Mutex{}
	t          trie.Trie
)

// Document is the struct that each item in the database
// is built with
type Document struct {
	ID   string
	Data interface{}
}

// Response struct formats responses from the database to external callers
type Response struct {
	Data    interface{}
	Success bool
	Meta    interface{}
}

// Collection is the struct that files are read into when opened
type Collection struct {
	Name string
	// Data []Document
	Data map[string]interface{}
	Meta interface{}
	Trie trie.Trie
}

// DB is the struct that the entire database uses
type DB struct {
	trie        trie.Trie
	Collections map[string]*Collection
}

// NewDocument creates a new document with a valid, unique UUID and
// and returns a Document or an error
func NewDocument(data interface{}) (Document, error) {
	uuid, err := newUUID()
	if err != nil {
		return Document{}, nil
	}

	return Document{
		ID:   uuid,
		Data: data,
	}, nil
}

// NewCollection creates a new Collection, creates the file, saves it,
// and returns a Collection or an error
func NewCollection(name string, trie trie.Trie) (Collection, error) {
	coll := Collection{
		Name: name,
		Meta: nil,
		Data: make(map[string]interface{}),
		Trie: trie,
	}

	err := SaveCollection(name, coll)
	if err != nil {
		log.Printf("error saving collection", err)
		return Collection{}, err
	}

	log.Printf("created new collection %+v\n", coll)

	return coll, nil
}

// GetCollection loads a collection if it exists. If it does not
// exist, it will create it and return that instead.
func GetCollection(name string) (Collection, error) {
	home, err := getPath()
	_, err = os.Stat(filepath.Join(home, name))

	if os.IsNotExist(err) {
		t := trie.New()
		coll, err := NewCollection(name, *t)
		return coll, err
	}
	return LoadCollection(name), nil
}

// getPath returns the .pangolin path
func getPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".pangolin")
	return path, nil
}

// SaveCollection is a helper function that takes a Collection
// and saves it. If the file doesn't exist, it will create it.
func SaveCollection(name string, coll Collection) error {
	home, err := getPath()
	path := filepath.Join(home, name)
	err = Save(path, coll)
	if err != nil {
		return Error.New("Error saving collection")
	}
	return err
}

// LoadCollection will read a file Collection into a Collection
// struct and returns that Collection
func LoadCollection(name string) Collection {
	home, err := getPath()
	checkError(err)
	path := filepath.Join(home, name)
	log.Printf("loading collection from path %s", path)
	coll := &Collection{}
	err = Load(path, coll)
	if err != nil {
		log.Fatal("Error loading collection: ", err)
	}
	return *coll
}

// SetupDatabase will create a new database and setup the config file
// if it does not exist
func SetupDatabase() (*DB, error) {
	home, err := PangolinHomeDir()
	checkError(err)
	createDirectory(filepath.Join(home))
	createFile(filepath.Join(home, ".config"))

	database := &DB{
		Collections: make(map[string]*Collection),
	}
	return database, nil
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

// InsertTrieKey puts a JSON blob into the collection
func (db *DB) InsertTrieKey(key string, data interface{}) (Response, error) {
	db.trie.Add(key, data)
	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// UpdateTrieKey inserts a value into the database
// If upsert is true, it will insert the data if the key is not found
func (db *DB) UpdateTrieKey(key string, data interface{}) (Response, error) {
	mutex.Lock()
	db.trie.Remove(key)
	db.trie.Add(key, data)
	mutex.Unlock()

	return Response{
		Data:    data,
		Success: true,
	}, nil
}

// DeleteTrieKey will delete an object from the tree
func (db *DB) DeleteTrieKey(key string) (Response, error) {
	mutex.Lock()
	db.trie.Remove(key)
	ok := db.trie.HasKeysWithPrefix(key)
	mutex.Unlock()
	if !ok {
		log.Printf("deleted key %s", key)
		return Response{
			Data:    key,
			Success: true,
		}, nil
	}
	log.Printf("could not delete key %s", key)

	return Response{
		Data:    nil,
		Success: false,
	}, Error.New("Delete Error: Key not deleted")
}

func (db *DB) getTrie() (trie.Trie, error) {
	return db.trie, nil
}

// PangolinHomeDir returns a string that is the home directory
// for Pangolin (defaults to `~/.pangolin`
func PangolinHomeDir() (string, error) {
	dir, err := homedir.Dir()
	f := filepath.Join(dir, ".pangolin")
	return f, err
}

func createFile(path string) {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Printf("can't create file %s - error: %s \n", file, err)
		}
		defer file.Close()
		checkError(err)
	}
}

func createDirectory(path string) {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var err = os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Printf("can't create directory %s", err)
		}
	}
}

// createCollectionFile creates the file for the collection
// and returns an error if it was unsuccessful.
func (db *DB) createCollectionFile(collection string) error {
	home, err := PangolinHomeDir()
	if err != nil {
		fmt.Printf("Error creating collection file: ", err)
		return err
	}
	collectionPath := filepath.Join(home, collection)
	createFile(collectionPath)
	return nil
}

// checkError is an error handler function that will exit the
// process if there are errors.
func checkError(err error) {
	if err != nil {
		fmt.Printf("ERROR", err.Error())
		os.Exit(0)
	}
}
