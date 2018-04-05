package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/petar/GoLLRB/llrb"
	"os"
	"runtime"
)

const DatabasePath = "/tmp/pangolin.db"

type DB struct {
	tree *llrb.LLRB
}

// interface needs work
type Database interface {
	Get (id int) 
	Set (blob Blob)
	Has (id int)
	Delete (id int)
	New (llrb.LLRB)
}

type Blob struct {
	id int `json:"id"`
	data string `json:"data"`
}

func (db *DB) Set (blob *Blob) llrb.Item {
	return db.tree.ReplaceOrInsert(blob)
}

func (db *DB) Get (id int) llrb.Item {
	blob := &Blob{id: id}
	return db.tree.Get(blob)
}

func (b *Blob) Less (than llrb.Item) bool {
	if v, ok := than.(*Blob); ok {
		return b.id < v.id
	}
	return false
}

func main() {
	fmt.Println("pangolin is starting up")

 	http.HandleFunc("/", getSpec)

	// tree := CreateTree() 

	// db := DB{tree}

	blob := &Blob{id: 1, data: "1234"}

	// db.Set(blob)

	// queryBlob := &Blob{id: 1}

	// fmt.Println(tree.Get(queryBlob))
	// fmt.Println(tree.Has(queryBlob))

	db, err := Load()
	Check(err)

	db.Set(blob)
	// writeErr := Save(file, tree)
	// Check(writeErr)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

type spec struct {
	Version string `json:"version"`
	Buckets []string `json:"buckets"`
}

func getBuckets () string { 
	return "none"
}

func NewSpec (version, buckets string) spec {
	return spec{
		Version: version,
		Buckets: []string{buckets},
	}
}

func getSpec (w http.ResponseWriter, r *http.Request) {
	response := NewSpec("0.0.1", getBuckets())
	message, _ := json.Marshal(response)
  w.Write([]byte(message))
}

func HandlePost (w http.ResponseWriter, r *http.Request) {

}

func HandleGet (w http.ResponseWriter, r *http.Request) {

}

func Save (path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func Load () error {
	if _, err := os.Stat(DatabasePath); os.IsNotExist(err) {
		fmt.Println("Database does not exist") // need to handle this
	}

	db := &DB{}

	file, err := os.Open(DatabasePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(db.tree)
	}
	file.Close()

	if db.tree == nil {
		db.tree = &llrb.LLRB{}
	}
	return err
}

func CreateTree () *llrb.LLRB {
	tree := llrb.New()
	return tree
}

func Load () (*DB, error) {
}
