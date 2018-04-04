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

const file = "/tmp/pangolin.db"

type Blob struct {
	id int `json:"id"`
	data string `json:"data"`
}

func (b *Blob) Less (than llrb.Item) bool {
	if v, ok := than.(*Blob); ok {
		return b.id < v.id
	}
	return false
}

// func (b *Blob) Less (than llrb.Item) bool {
// 	switch than {
// 	case nil: 
// 		return false
// 	default:
// 		return b.id < Blob(than).id
// 	}
// }

// func lessInt(a, b interface{}) bool { return a.(int) < b.(int) }

func main() {
	fmt.Println("pangolin is starting up")

 	http.HandleFunc("/", getSpec)

	tree := CreateTree() 

	blob := &Blob{id: 1, data: "1234"}

	tree.ReplaceOrInsert(blob)

	queryBlob := &Blob{id: 1}

	fmt.Println(tree.Get(queryBlob))

	writeErr := Save(file, tree)
	Check(writeErr)

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

func getSpec (w http.ResponseWriter, r *http.Request) {
	response := &spec{
		Version: "0.0.1",
		Buckets: []string{getBuckets()},
	}

	message, _ := json.Marshal(response)
  w.Write([]byte(message))
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

func Load (path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func CreateTree () *llrb.LLRB {
	tree := llrb.New()
	return tree
}
